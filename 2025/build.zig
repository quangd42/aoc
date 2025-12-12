const std = @import("std");

pub fn build(b: *std.Build) void {
    const target = b.standardTargetOptions(.{});
    const optimize = b.standardOptimizeOption(.{});

    const day_option = b.option(u8, "day", "Advent of Code day to compile") orelse 0;

    var day_buf: [2]u8 = undefined;
    const day_exe_name = std.fmt.bufPrint(&day_buf, "{d:0>2}", .{day_option}) catch @panic("OOM");

    var path_buf: [64]u8 = undefined;
    const path_string = std.fmt.bufPrint(&path_buf, "{s}/main.cpp", .{day_exe_name}) catch @panic("OOM");

    var cwd = std.fs.cwd();
    cwd.access(path_string, .{ .mode = .read_only }) catch |err| switch (err) {
        error.FileNotFound => {
            std.debug.print("FileNotFound: {s}\n", .{path_string});
            std.process.exit(1);
        },
        else => |e| {
            std.debug.print("{}\n", .{e});
            std.process.exit(1);
        },
    };

    const exe = b.addExecutable(.{
        .name = day_exe_name,
        .root_module = b.createModule(.{
            .target = target,
            .optimize = optimize,
            .link_libcpp = true,
        }),
    });
    exe.root_module.addCSourceFile(.{ .file = b.path(path_string), .language = .cpp });

    b.installArtifact(exe);

    const run_step = b.step("run", "Run the solution");

    const run_cmd = b.addRunArtifact(exe);
    run_step.dependOn(&run_cmd.step);

    run_cmd.step.dependOn(b.getInstallStep());

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
