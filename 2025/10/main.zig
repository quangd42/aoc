const std = @import("std");
const mem = std.mem;

const MAX_N = 10;
const MAX_BUTTONS = 16;
const STATE_CAP = 1 << MAX_N;

const Parsed = struct {
    goal: u16 = 0, // bit i is 1 if goal[i] == '#'
    masks: [MAX_BUTTONS]u16 = [_]u16{0} ** MAX_BUTTONS,
    nb: u8 = 0, // number of buttons
    n: u8 = 0, // goal length (<= 10)
};

fn parse_machine(line: []const u8) !Parsed {
    var p: Parsed = .{};

    // Find goal between [ ... ]
    const l = mem.indexOfScalar(u8, line, '[') orelse return error.InvalidLine;
    const r = mem.indexOfScalarPos(u8, line, l + 1, ']') orelse return error.InvalidLine;

    if (r <= l + 1) return error.InvalidLine;

    const len = r - l - 1;
    if (len > MAX_N) return error.GoalTooLong;

    p.n = @intCast(len);

    for (line[l + 1 .. r], 0..) |c, i| {
        switch (c) {
            '#' => p.goal |= @as(u16, 1) << @intCast(i),
            '.' => {},
            else => return error.InvalidGoalChar,
        }
    }

    // Parse buttons "(...)" after the goal
    {
        var end: usize = 0;
        while (mem.indexOfScalarPos(u8, line, end + 1, '(')) |start| {
            end = mem.indexOfScalarPos(u8, line, start, ')') orelse return error.InvalidLine;

            if (p.nb >= MAX_BUTTONS) return error.TooManyButtons;

            var mask: u16 = 0;
            var it = mem.splitScalar(u8, line[start + 1 .. end], ',');
            while (it.next()) |num| {
                const idx = try std.fmt.parseUnsigned(u8, num, 10);
                if (idx >= p.n) return error.ButtonIndexOutOfRange;
                mask ^= @as(u16, 1) << @intCast(idx);
            }

            p.masks[p.nb] = mask;
            p.nb += 1;
        }
    }

    return p;
}

fn bfs(p: *const Parsed) u16 {
    var distance = [_]?u16{null} ** STATE_CAP;
    distance[0] = 0;

    var q: [STATE_CAP]u16 = undefined;
    var head: usize = 0;
    var tail: usize = 0;
    q[tail] = 0;
    tail += 1;

    while (head != tail) {
        const s = q[head];
        head += 1;
        if (s == p.goal) return @intCast(distance[s].?);
        for (0..p.nb) |i| {
            const t = s ^ p.masks[i];
            if (distance[t] == null) {
                distance[t] = distance[s].? + 1;
                q[tail] = t;
                tail += 1;
            }
        }
    }
    unreachable;
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
    while (reader.takeDelimiterInclusive('\n')) |line| {
        var p = try parse_machine(line);
        sum += bfs(&p);
    } else |err| {
        if (err != error.EndOfStream) {
            std.debug.print("{}", .{err});
        }
    }
    std.debug.print("{d}\n", .{sum});
}
