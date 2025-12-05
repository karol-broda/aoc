#include <stdio.h>
#include <stdlib.h>
#include <string.h>

#define MAX_LINES 10000
#define MAX_LINE_LEN 256

static int part1(char lines[][MAX_LINE_LEN], int line_count);
static int part2(char lines[][MAX_LINE_LEN], int line_count);

int main(int argc, char *argv[]) {
    FILE *fp = NULL;
    char lines[MAX_LINES][MAX_LINE_LEN];
    int line_count = 0;

    const char *filename = (argc > 1) ? argv[1] : "input.txt";
    fp = fopen(filename, "r");
    if (fp == NULL) {
        fprintf(stderr, "error: cannot open %s\n", filename);
        return 1;
    }

    while (fgets(lines[line_count], MAX_LINE_LEN, fp) != NULL && line_count < MAX_LINES) {
        size_t len = strlen(lines[line_count]);
        if (len > 0 && lines[line_count][len - 1] == '\n') {
            lines[line_count][len - 1] = '\0';
        }
        line_count++;
    }
    fclose(fp);

    printf("part 1: %d\n", part1(lines, line_count));
    printf("part 2: %d\n", part2(lines, line_count));

    return 0;
}

static int part1(char lines[][MAX_LINE_LEN], int line_count) {
    int position = 50;
    int count = 0;

    for (int i = 0; i < line_count; i++) {
        if (strlen(lines[i]) == 0) {
            continue;
        }

        char direction = lines[i][0];
        int distance = atoi(&lines[i][1]);

        if (direction == 'L') {
            position = (position - distance) % 100;
            if (position < 0) {
                position = position + 100;
            }
        } else if (direction == 'R') {
            position = (position + distance) % 100;
        }

        if (position == 0) {
            count++;
        }
    }

    return count;
}

static int part2(char lines[][MAX_LINE_LEN], int line_count) {
    int position = 50;
    int count = 0;

    for (int i = 0; i < line_count; i++) {
        if (strlen(lines[i]) == 0) {
            continue;
        }

        char direction = lines[i][0];
        int distance = atoi(&lines[i][1]);

        if (direction == 'L') {
            if (position == 0) {
                count += distance / 100;
            } else {
                count += distance / 100;
                if (distance % 100 >= position) {
                    count++;
                }
            }
            position = (position - distance) % 100;
            if (position < 0) {
                position = position + 100;
            }
        } else if (direction == 'R') {
            count += (position + distance) / 100;
            position = (position + distance) % 100;
        }
    }

    return count;
}

