const std = @import("std");
const mem = std.mem;

const AdjMatrix = std.StringArrayHashMap([][]const u8);

pub fn parse(am: *AdjMatrix, line: []const u8) !void {
    const colon = mem.indexOfScalar(u8, line, ':') orelse return error.InvalidLine;
    const vertex = line[0..colon];
    var it = mem.tokenizeScalar(u8, line[colon + 1 ..], ' ');
    const gpa = am.allocator;
    var neighbors: std.ArrayList([]const u8) = .empty;
    while (it.next()) |n| {
        try neighbors.append(gpa, try gpa.dupe(u8, n));
    }

    try am.put(try gpa.dupe(u8, vertex), try neighbors.toOwnedSlice(gpa));
}

pub fn print_map(am: *AdjMatrix) void {
    for (am.keys()) |key| {
        std.debug.print("{s}: ", .{key});
        const values = am.get(key) orelse {
            std.debug.print("key not found: {s}\n", .{key});
            return;
        };
        for (values) |n| {
            std.debug.print("{s} ", .{n});
        }
        std.debug.print("\n", .{});
    }
}

pub fn dfs_part1(am: *AdjMatrix) !u64 {
    const gpa = am.allocator;
    var stack: std.ArrayList([]const u8) = .empty;
    try stack.append(gpa, "you");
    // var visited = std.StringArrayHashMap(bool).init(gpa);
    var out: u64 = 0;
    while (stack.items.len > 0) {
        const v = stack.pop().?;
        if (mem.eql(u8, v, "out")) {
            out += 1;
            continue;
        }

        // const gop = try visited.getOrPut(v);
        // if (gop.found_existing) continue;
        // gop.value_ptr.* = true;

        const neighbors = am.get(v).?;
        try stack.appendSlice(gpa, neighbors);
    }
    return out;
}

fn is_valid_path(path: [][]const u8) bool {
    var dac = false;
    var fft = false;
    for (path) |l| {
        if (mem.eql(u8, "dac", l)) dac = true;
        if (mem.eql(u8, "fft", l)) fft = true;
    }
    return dac and fft;
}

pub fn main() !void {
    var arena = std.heap.ArenaAllocator.init(std.heap.page_allocator);
    defer arena.deinit();
    const alloc = arena.allocator();

    const argv = try std.process.argsAlloc(alloc);
    if (argv.len < 2) {
        std.debug.print("missing input argv\n", .{});
        return;
    }
    const file = try std.fs.openFileAbsolute(argv[1], .{ .mode = .read_only });
    defer file.close();

    var buffer: [1024]u8 = undefined;
    var file_reader = file.reader(&buffer);
    const reader = &file_reader.interface;

    var am = AdjMatrix.init(alloc);

    var sum: u64 = 0;
    while (reader.takeDelimiterExclusive('\n')) |line| : (reader.toss(1)) {
        // process line here
        try parse(&am, line);
    } else |err| switch (err) {
        error.EndOfStream => {},
        else => std.debug.print("{}", .{err}),
    }

    sum = try dfs_part1(&am);
    std.debug.print("{d}\n", .{sum});
}
