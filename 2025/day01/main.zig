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
    var position: i32 = 50;
    var count: i32 = 0;

    var lines = std.mem.splitScalar(u8, input, '\n');
    while (lines.next()) |line| {
        if (line.len == 0) continue;

        const direction = line[0];
        const distance = std.fmt.parseInt(i32, line[1..], 10) catch continue;

        // rotate the dial, L = left (counterclockwise), R = right (clockwise)
        if (direction == 'L') {
            position = @mod(position - distance, 100);
        } else if (direction == 'R') {
            position = @mod(position + distance, 100);
        }

        if (position == 0) {
            count += 1;
        }
    }

    return count;
}

fn part2(input: []const u8) i32 {
    var position: i32 = 50;
    var count: i32 = 0;

    var lines = std.mem.splitScalar(u8, input, '\n');
    while (lines.next()) |line| {
        if (line.len == 0) continue;

        const direction = line[0];
        const distance = std.fmt.parseInt(i32, line[1..], 10) catch continue;

        // count how many times the dial passes through 0 during rotation
        if (direction == 'L') {
            if (position == 0) {
                count += @divTrunc(distance, 100);
            } else {
                count += @divTrunc(distance, 100);
                if (@mod(distance, 100) >= position) {
                    count += 1;
                }
            }
            position = @mod(position - distance, 100);
        } else if (direction == 'R') {
            count += @divTrunc(position + distance, 100);
            position = @mod(position + distance, 100);
        }
    }

    return count;
}
