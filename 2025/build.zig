const std = @import("std");

const input = enum {
    example,
    full,

    pub fn path(i: input, day: u8, buf: []u8) []const u8 {
        var path_buf: [64]u8 = undefined;
        const cwd_path = std.fs.cwd().realpath(".", &path_buf) catch @panic("failed to get cwd's real path");
        return switch (i) {
            .example => std.fmt.bufPrint(buf, "{s}/{d:0>2}/example.txt", .{ cwd_path, day }) catch @panic("OOM"),
            .full => std.fmt.bufPrint(buf, "{s}/{d:0>2}/input.txt", .{ cwd_path, day }) catch @panic("OOM"),
        };
    }
};

pub fn build(b: *std.Build) !void {
    const target = b.standardTargetOptions(.{});
    const optimize = b.standardOptimizeOption(.{});

    const day_option = b.option(u8, "day", "Advent of Code day to compile") orelse 1;
    const input_option = b.option(input, "input", "Input file") orelse .example;

    var path_buf: [64]u8 = undefined;
    const source_path = std.fmt.bufPrint(&path_buf, "{d:0>2}/main.cpp", .{day_option}) catch @panic("OOM");

    var cwd = std.fs.cwd();
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

    const exe = b.addExecutable(.{
        .name = "aoc_2025",
        .root_module = b.createModule(.{
            .target = target,
            .optimize = optimize,
            .link_libcpp = true,
        }),
    });
    exe.root_module.addCSourceFile(.{ .file = b.path(source_path), .language = .cpp });

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

    // Just like flags, top level steps are also listed in the `--help` menu.
    //
    // The Zig build system is entirely implemented in userland, which means
    // that it cannot hook into private compiler APIs. All compilation work
    // orchestrated by the build system will result in other Zig compiler
    // subcommands being invoked with the right flags defined. You can observe
    // these invocations when one fails (or you pass a flag to increase
    // verbosity) to validate assumptions and diagnose problems.
    //
    // Lastly, the Zig build system is relatively simple and self-contained,
    // and reading its source code will allow you to master it.
}
