const std = @import("std");
const mem = std.mem;
const Allocator = mem.Allocator;
const ArenaAllocator = std.heap.ArenaAllocator;
const ArrayList = std.ArrayList;

const Button = []const usize;
const Lights = []bool;

const Machine = struct {
    goal: Lights,
    buttons: []Button,
    backing_allocator: Allocator,
    arena: ArenaAllocator,

    fn init(backing: Allocator, line: []const u8) !Machine {
        var arena = std.heap.ArenaAllocator.init(backing);
        const gpa = arena.allocator();
        var goal: ArrayList(bool) = .empty;
        {
            const start = mem.indexOfScalar(u8, line, '[').?;
            const end = mem.indexOfScalar(u8, line, ']').?;
            try goal.ensureTotalCapacity(gpa, end - start);
            for (line[start + 1 .. end]) |c| {
                switch (c) {
                    '.' => goal.appendAssumeCapacity(false),
                    '#' => goal.appendAssumeCapacity(true),
                    else => return error.InvalidInput,
                }
            }
        }

        var buttons: std.ArrayList(Button) = .empty;
        {
            var end: usize = 0;
            var start: usize = 0;
            while (true) {
                start = mem.indexOfScalarPos(u8, line, end + 1, '(') orelse break;
                end = mem.indexOfScalarPos(u8, line, start, ')').?;
                var it = mem.splitScalar(u8, line[start + 1 .. end], ',');
                var button: ArrayList(usize) = .empty;
                while (it.next()) |bs| {
                    try button.append(gpa, try std.fmt.parseUnsigned(usize, bs, 10));
                }
                try buttons.append(gpa, try button.toOwnedSlice(gpa));
            }
        }

        return .{
            .goal = try goal.toOwnedSlice(gpa),
            .buttons = try buttons.toOwnedSlice(gpa),
            .backing_allocator = backing,
            .arena = arena,
        };
    }

    fn deinit(self: *Machine) void {
        self.arena.deinit();
    }
};

const State = struct {
    prev: ?*State = null,
    current: Lights,
    next: ?[]State = null,
    backing_allocator: Allocator,
    arena: ArenaAllocator,

    fn init(backing: Allocator, current: Lights, prev: ?*State) State {
        return .{
            .backing_allocator = backing,
            .arena = ArenaAllocator.init(backing),
            .current = current,
            .prev = prev,
        };
    }

    fn initRoot(backing: Allocator, light_size: usize) !State {
        var arena = ArenaAllocator.init(backing);
        const gpa = arena.allocator();
        const current = try gpa.alloc(bool, light_size);
        return .{
            .backing_allocator = backing,
            .arena = arena,
            .current = current,
            .prev = null,
        };
    }

    fn deinit(self: *State) void {
        self.arena.deinit();
    }

    fn populate_next(self: *State, buttons: []Button) !void {
        const gpa = self.arena.allocator();
        if (self.next) |_| return;
        var next: std.ArrayList(State) = .empty;
        try next.ensureTotalCapacity(gpa, buttons.len);

        for (buttons) |button| {
            var new_mstate = try gpa.dupe(bool, self.current);
            for (button) |i| {
                new_mstate[i] = !new_mstate[i];
            }
            next.appendAssumeCapacity(State.init(self.backing_allocator, new_mstate, self));
        }

        self.next = try next.toOwnedSlice(gpa);
    }
};

fn bfs_state(root: *State, m: *Machine) !u64 {
    var q: [1 << 10]*State = undefined;
    var head: usize = 0;
    var tail: usize = 0;
    q[tail] = root;
    tail += 1;

    while (head != tail) {
        const state: *State = q[head];
        head += 1;
        if (mem.eql(bool, state.current, m.goal)) {
            var dist: usize = 0;
            var cur = state;
            while (cur.prev) |prev| {
                cur = prev;
                dist += 1;
            }
            return dist;
        }

        try state.populate_next(m.buttons);
        for (state.next.?) |*branch| {
            q[tail] = branch;
            tail += 1;
        }
    }

    std.debug.print("bfs search queue is empty or full\n", .{});
    std.process.abort();
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
        var m = try Machine.init(alloc, line);
        var root = try State.initRoot(alloc, m.goal.len);
        sum += try bfs_state(&root, &m);
    } else |err| {
        if (err != error.EndOfStream) {
            std.debug.print("{}", .{err});
        }
    }
    std.debug.print("{d}\n", .{sum});
}
