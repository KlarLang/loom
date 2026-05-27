const std = @import("std");

pub fn commandEqual(command: []const u8, potential_command: []const u8) bool {
    if (command.len != potential_command.len) return false;

    for (command, potential_command) |c, pc| {
        if (c != pc) return false;
    }

    return true;
}

pub fn flagsEqual(flag: []const u8, potential_flags: []const []const u8) bool {
    for (potential_flags) |pf| {
        if (commandEqual(flag, pf)) return true;
    }

    return false;
}
