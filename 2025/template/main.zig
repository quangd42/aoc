const std = @import("std");

pub fn process_line(line: []const u8) u64 {
    _ = line;
    return 0;
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

    var sum: u64 = 0;
    while (reader.takeDelimiterExclusive('\n')) |line| : (reader.toss(1)) {
        // process line here
        sum += process_line(line);
    } else |err| switch (err) {
        error.EndOfStream => {},
        else => std.debug.print("{}", .{err}),
    }

    std.debug.print("{d}\n", .{sum});
}
