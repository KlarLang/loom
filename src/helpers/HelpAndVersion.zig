const builtin = @import("build_options");
const IoWriter = @import("IoWriter.zig");

pub fn help(writer: IoWriter) !void {
    try writer.stdout.interface.print("Usage: loom [file|file.fiber]\n\n", .{});

    try writer.stdout.interface.print("Flags:\n", .{});
    try writer.stdout.interface.print("  -h, --help      Show this help.\n", .{});
    try writer.stdout.interface.print("  -v, --version   Show loom version.\n\n", .{});

    try writer.stdout.flush();
}

pub fn version(writer: IoWriter) !void {
    try writer.stdout.interface.print("loom {s}\n", .{builtin.loom_version});
    try writer.stdout.flush();
}
