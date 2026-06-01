const std = @import("std");
const Opcodes = @import("opcodes.zig").Opcodes;

fn Stack(comptime T: type) type {
    return struct {
        list: std.ArrayList(T),

        const Self = @This();

        pub fn init() Self {
            return .{
                .list = .empty,
            };
        }

        pub fn push(self: *Self, alloc: std.mem.Allocator, value: T) !void {
            try self.list.append(alloc, value);
        }

        pub fn pop(self: *Self) !T {
            return self.list.pop();
        }

        pub fn deinit(self: *Self, alloc: std.mem.Allocator) !void {
            self.list.deinit(alloc);
        }
    };
}

pub const VM = struct {
    global_stack: []const u8,
    file_path: []const u8,

    fn entry(_: []const u8) void {
        while (true) {
            // fetch
            const test_ = Opcodes.PUSH_F16;

            switch (test_) {}
        }
    }
};
