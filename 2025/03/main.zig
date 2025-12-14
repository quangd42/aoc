const std = @import("std");

pub fn largestn(s: []u8, idx: usize, buf: []u8) void {
    if (idx < 1 or s.len < idx) return;
    var hi: u8 = s[0];
    var hi_i: usize = 0;

    for (s[0 .. s.len - idx], 0..) |c, i| {
        if (c > hi) {
            hi = c;
            hi_i = i;
        }
    }
    buf[0] = hi;
    largestn(s[hi_i + 1 ..], idx - 1, buf[1..]);
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
    var num_buf: [64]u8 = undefined;
    while (reader.takeDelimiterInclusive('\n')) |line| {
        largestn(line, 12, &num_buf);
        sum += try std.fmt.parseInt(u64, num_buf[0..12], 10);
        // std.debug.print("{s}\n", .{num_buf[0..2]});
    } else |err| {
        if (err != error.EndOfStream) {
            std.debug.print("{}", .{err});
        }
    }
    std.debug.print("{d}\n", .{sum});
}
