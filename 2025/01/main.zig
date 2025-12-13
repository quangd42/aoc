const std = @import("std");

pub fn main() !void {
    // var arena = std.heap.ArenaAllocator.init(std.heap.page_allocator);
    // const a = arena.allocator();

    const file = try std.fs.cwd().openFile("input.txt", .{});
    defer file.close();

    var buffer: [1024]u8 = undefined;
    var file_reader = file.reader(&buffer);
    const reader = &file_reader.interface;

    var lock = Lock.init();
    while (reader.takeDelimiterInclusive('\n')) |line| {
        const inst = try parse_inst(line[0 .. line.len - 1]);
        try lock.dial2(inst);
    } else |err| {
        if (err != error.EndOfStream) {
            std.debug.print("got error {}", .{err});
        }
    }
    std.debug.print("{d}\n", .{lock.count});
}

const Inst = struct {
    lr: bool,
    count: i16,
};

const Lock = struct {
    state: i16 = 50,
    count: u16,

    fn init() Lock {
        return .{ .count = 0 };
    }

    fn dial(self: *Lock, inst: Inst) !void {
        if (!inst.lr) {
            self.state -= inst.count;
        } else {
            self.state += inst.count;
        }

        self.state = @mod(self.state, 100);
        if (self.state == 0) {
            self.count += 1;
        }
    }

    fn dial2(self: *Lock, inst: Inst) !void {
        var count = inst.count;
        while (count > 0) {
            count -= 1;
            if (!inst.lr) {
                self.state -= 1;
            } else {
                self.state += 1;
            }
            self.state = @mod(self.state, 100);
            if (self.state == 0) {
                self.count += 1;
            }
        }
    }
};

fn parse_inst(line: []u8) !Inst {
    var lr = false;
    switch (line[0]) {
        'L' => {},
        'R' => lr = true,
        else => unreachable,
    }
    const count = try std.fmt.parseInt(i16, line[1..], 10);
    return .{ .lr = lr, .count = count };
}
