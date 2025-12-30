const std = @import("std");

pub fn build(b: *std.Build) !void {
    const target = b.standardTargetOptions(.{});
    const optimize = b.standardOptimizeOption(.{});

    var day_option = b.option(u8, "day", "Advent of Code day to compile") orelse 0;
    const input_option = b.option(input, "input", "Input file") orelse .example;
    const lang_option = b.option(lang, "lang", "Language the solution is in") orelse .cpp;

    var cwd = std.fs.cwd();
    if (day_option == 0) {
        var cur_dir = try cwd.openDir(".", .{ .iterate = true });
        defer cur_dir.close();

        var it = cur_dir.iterate();
        var hi: u8 = 0;
        while (try it.next()) |entry| {
            switch (entry.kind) {
                .directory => {
                    const n = std.fmt.parseInt(u8, entry.name, 10) catch continue;
                    if (n > hi) hi = n;
                },
                else => {},
            }
        }
        if (hi == 0) {
            std.debug.print("No solution has been created yet!", .{});
            std.process.exit(0);
        }
        day_option = hi;
    }

    var path_buf: [64]u8 = undefined;
    const source_path = std.fmt.bufPrint(
        &path_buf,
        "{d:0>2}/main.{s}",
        .{ day_option, lang_option.file_ext() },
    ) catch @panic("OOM");

    cwd.access(source_path, .{ .mode = .read_only }) catch |err| switch (err) {
        error.FileNotFound => {
            std.debug.print("FileNotFound: {s}\n", .{source_path});
            std.process.exit(1);
        },
        else => |e| {
            std.debug.print("{}\n", .{e});
            std.process.exit(1);
        },
    };

    const exe_mod = blk: switch (lang_option) {
        .cpp => {
            var mod = b.createModule(.{
                .target = target,
                .optimize = optimize,
                .link_libcpp = true,
            });
            mod.addCSourceFile(.{ .file = b.path(source_path), .language = .cpp });
            break :blk mod;
        },
        .zig => b.createModule(.{
            .target = target,
            .optimize = optimize,
            .root_source_file = b.path(source_path),
        }),
    };

    const exe = b.addExecutable(.{
        .name = "aoc_2025",
        .root_module = exe_mod,
    });

    b.installArtifact(exe);

    const run_step = b.step("run", "Run the solution");

    const run_cmd = b.addRunArtifact(exe);
    run_step.dependOn(&run_cmd.step);

    run_cmd.step.dependOn(b.getInstallStep());

    var arg_buf: [64]u8 = undefined;
    run_cmd.addArg(input_option.path(day_option, &arg_buf));

    if (b.args) |args| {
        run_cmd.addArgs(args);
    }
}

const input = enum {
    example,
    full,

    pub fn path(i: input, day: u8, buf: []u8) []const u8 {
        var path_buf: [64]u8 = undefined;
        const cwd_path = std.fs.cwd().realpath(".", &path_buf) catch @panic("failed to get cwd's real path");
        return switch (i) {
            .example => std.fmt.bufPrint(buf, "{s}/{d:0>2}/example.txt", .{ cwd_path, day }) catch @panic("OOM"),
            .full => std.fmt.bufPrint(buf, "{s}/{d:0>2}/full.txt", .{ cwd_path, day }) catch @panic("OOM"),
        };
    }
};

const lang = enum {
    cpp,
    zig,

    fn file_ext(l: lang) []const u8 {
        return switch (l) {
            .cpp => "cpp",
            .zig => "zig",
        };
    }
};
