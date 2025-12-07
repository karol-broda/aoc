const std = @import("std");

pub fn main() !void {
    var gpa = std.heap.GeneralPurposeAllocator(.{}){};
    defer _ = gpa.deinit();
    const allocator = gpa.allocator();

    const input = try readInput(allocator);
    defer allocator.free(input);

    const stdout = std.io.getStdOut().writer();
    try stdout.print("part 1: {d}\n", .{part1(input)});
    try stdout.print("part 2: {d}\n", .{part2(input)});
}

fn readInput(allocator: std.mem.Allocator) ![]const u8 {
    const stdin = std.io.getStdIn();
    const stat = stdin.stat() catch {
        return readFile(allocator, "input.txt");
    };

    if (stat.kind == .character_device) {
        return readFile(allocator, "input.txt");
    }

    return try stdin.readToEndAlloc(allocator, std.math.maxInt(usize));
}

fn readFile(allocator: std.mem.Allocator, filename: []const u8) ![]const u8 {
    const file = std.fs.cwd().openFile(filename, .{}) catch |err| {
        std.debug.print("error: cannot open {s}\n", .{filename});
        return err;
    };
    defer file.close();
    return try file.readToEndAlloc(allocator, std.math.maxInt(usize));
}

fn part1(input: []const u8) i32 {
    _ = input;
    return 0;
}

fn part2(input: []const u8) i32 {
    _ = input;
    return 0;
}
