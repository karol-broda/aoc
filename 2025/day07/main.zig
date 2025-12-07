const std = @import("std");

pub fn main() !void {
    var gpa = std.heap.GeneralPurposeAllocator(.{}){};
    defer _ = gpa.deinit();
    const allocator = gpa.allocator();

    const input = try readInput(allocator);
    defer allocator.free(input);

    const stdout = std.io.getStdOut().writer();
    try stdout.print("part 1: {d}\n", .{try part1(allocator, input)});
    try stdout.print("part 2: {d}\n", .{try part2(allocator, input)});
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

fn part1(allocator: std.mem.Allocator, input: []const u8) !u64 {
    var lines = std.ArrayList([]const u8).init(allocator);
    defer lines.deinit();

    var line_iter = std.mem.splitScalar(u8, input, '\n');
    while (line_iter.next()) |line| {
        if (line.len > 0) {
            try lines.append(line);
        }
    }

    if (lines.items.len == 0) return 0;

    var start_col: ?usize = null;
    for (lines.items[0], 0..) |c, col| {
        if (c == 'S') {
            start_col = col;
            break;
        }
    }

    if (start_col == null) return 0;

    // track beam positions as a set; multiple beams at same column merge into one
    var beams = std.AutoHashMap(usize, void).init(allocator);
    defer beams.deinit();
    try beams.put(start_col.?, {});

    var split_count: u64 = 0;

    for (lines.items[1..]) |row| {
        var new_beams = std.AutoHashMap(usize, void).init(allocator);

        var beam_iter = beams.keyIterator();
        while (beam_iter.next()) |col_ptr| {
            const col = col_ptr.*;
            if (col < row.len) {
                const cell = row[col];
                if (cell == '^') {
                    // splitter stops beam, emits two new beams left and right
                    split_count += 1;
                    if (col > 0) {
                        try new_beams.put(col - 1, {});
                    }
                    if (col + 1 < row.len) {
                        try new_beams.put(col + 1, {});
                    }
                } else {
                    try new_beams.put(col, {});
                }
            }
        }

        beams.deinit();
        beams = new_beams;
    }

    return split_count;
}

fn part2(allocator: std.mem.Allocator, input: []const u8) !u64 {
    var lines = std.ArrayList([]const u8).init(allocator);
    defer lines.deinit();

    var line_iter = std.mem.splitScalar(u8, input, '\n');
    while (line_iter.next()) |line| {
        if (line.len > 0) {
            try lines.append(line);
        }
    }

    if (lines.items.len == 0) return 0;

    var start_col: ?usize = null;
    for (lines.items[0], 0..) |c, col| {
        if (c == 'S') {
            start_col = col;
            break;
        }
    }

    if (start_col == null) return 0;

    // track number of timelines at each column; timelines don't merge
    var particles = std.AutoHashMap(usize, u64).init(allocator);
    defer particles.deinit();
    try particles.put(start_col.?, 1);

    for (lines.items[1..]) |row| {
        var new_particles = std.AutoHashMap(usize, u64).init(allocator);

        var iter = particles.iterator();
        while (iter.next()) |entry| {
            const col = entry.key_ptr.*;
            const count = entry.value_ptr.*;

            if (col < row.len) {
                if (row[col] == '^') {
                    // each timeline splits into two
                    if (col > 0) {
                        const prev = new_particles.get(col - 1) orelse 0;
                        try new_particles.put(col - 1, prev + count);
                    }
                    if (col + 1 < row.len) {
                        const prev = new_particles.get(col + 1) orelse 0;
                        try new_particles.put(col + 1, prev + count);
                    }
                } else {
                    const prev = new_particles.get(col) orelse 0;
                    try new_particles.put(col, prev + count);
                }
            }
        }

        particles.deinit();
        particles = new_particles;
    }

    var total: u64 = 0;
    var value_iter = particles.valueIterator();
    while (value_iter.next()) |count| {
        total += count.*;
    }

    return total;
}
