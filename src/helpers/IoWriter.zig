const std = @import("std");

pub const IoWriter = @This();

sys: std.process.Init,

stdout: *std.Io.File.Writer,
stderr: *std.Io.File.Writer,
