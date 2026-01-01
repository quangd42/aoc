const std = @import("std");
const mem = std.mem;
const Allocator = mem.Allocator;

const Shape = []const u8;

pub fn parse_shapes(gpa: Allocator, block: []const u8) !Shape {
    const colon = mem.indexOfScalar(u8, block, ':') orelse return error.InvalidInput;
    _ = std.fmt.parseUnsigned(u8, block[0..colon], 10) catch return error.NotAShape;

    const idx = mem.indexOfScalar(u8, block, '\n') orelse return error.InvalidInput;
    return gpa.dupe(u8, block[idx + 1 ..]);
}

pub fn count_tiles(s: Shape) u8 {
    var out: u8 = 0;
    for (s) |c| {
        switch (c) {
            '#' => out += 1,
            else => {},
        }
    }
    return out;
}

const Region = struct {
    w: u32,
    h: u32,
    quantity: []u32,
    pub fn format(
        self: @This(),
        writer: *std.Io.Writer,
    ) std.Io.Writer.Error!void {
        try writer.print("{d}x{d}: ", .{ self.w, self.h });
        for (self.quantity) |q| {
            try writer.print("{d} ", .{q});
        }
    }
};

pub fn parse_region(gpa: Allocator, line: []const u8) !Region {
    const colon = mem.indexOfScalar(u8, line, ':') orelse return error.InvalidInput;

    const mult = mem.indexOfScalar(u8, line, 'x') orelse return error.InvalidInput;
    const w = std.fmt.parseUnsigned(u32, line[0..mult], 10) catch return error.InvalidInput;
    const h = std.fmt.parseUnsigned(u32, line[mult + 1 .. colon], 10) catch return error.InvalidInput;

    var it = mem.tokenizeScalar(u8, line[colon + 1 ..], ' ');
    const quantity = try gpa.alloc(u32, 6);
    var q_idx: usize = 0;
    while (it.next()) |n| : (q_idx += 1) {
        const q = std.fmt.parseUnsigned(u32, n, 10) catch return error.InvalidInput;
        quantity[q_idx] = q;
    }
    // std.debug.print("parsed quantity: ", .{});
    // for (quantity) |q| {
    //     std.debug.print("{d} ", .{q});
    // }
    // std.debug.print("\n", .{});
    return .{ .w = w, .h = h, .quantity = quantity };
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

    var shapes: [6]Shape = undefined;
    var regions: std.ArrayList(Region) = .empty;

    var shapes_idx: usize = 0;
    while (shapes_idx < 6) {
        _ = try reader.discardDelimiterInclusive('\n');
        shapes[shapes_idx] = try alloc.dupe(u8, try reader.take(13));
        shapes_idx += 1;
    }

    for (shapes, 0..) |s, i| {
        std.debug.print("{d}:\n{s}", .{ i, s });
    }

    while (reader.takeDelimiterExclusive('\n')) |line| : (reader.toss(1)) {
        try regions.append(alloc, try parse_region(alloc, line));
    } else |err| switch (err) {
        error.EndOfStream => {},
        else => return err,
    }

    var tile_count: [6]u8 = undefined;
    for (shapes, 0..) |s, i| {
        tile_count[i] = count_tiles(s);
    }

    var sum: u64 = 0;

    for (regions.items) |r| {
        std.debug.print("{f}\n", .{r});
        const avail: u32 = r.h * r.w;
        var required: u32 = 0;
        std.debug.print(" - avail: w*h = {d}\n", .{avail});
        std.debug.print(" - required: ", .{});
        for (r.quantity, 0..) |q, i| {
            const t = tile_count[i];
            required += q * t;
            std.debug.print("{d} * {d} + ", .{ q, t });
        }
        std.debug.print("= {d}\n", .{required});
        if (required <= avail) sum += 1;
    }

    std.debug.print("{d}\n", .{sum});
}
