#include <stdio.h>
#include <stdlib.h>
#include <string.h>

static long part1(const char *input, size_t len);
static long part2(const char *input, size_t len);

static char *read_file(const char *filename, size_t *out_len) {
    FILE *fp = fopen(filename, "rb");
    if (fp == NULL) {
        return NULL;
    }

    fseek(fp, 0, SEEK_END);
    long size = ftell(fp);
    fseek(fp, 0, SEEK_SET);

    char *buf = malloc(size + 1);
    if (buf == NULL) {
        fclose(fp);
        return NULL;
    }

    size_t read = fread(buf, 1, size, fp);
    buf[read] = '\0';
    fclose(fp);

    if (out_len != NULL) {
        *out_len = read;
    }
    return buf;
}

int main(int argc, char *argv[]) {
    const char *filename = (argc > 1) ? argv[1] : "input.txt";

    size_t len = 0;
    char *input = read_file(filename, &len);
    if (input == NULL) {
        fprintf(stderr, "error: cannot read %s\n", filename);
        return 1;
    }

    printf("part 1: %ld\n", part1(input, len));
    printf("part 2: %ld\n", part2(input, len));

    free(input);
    return 0;
}

static long part1(const char *input, size_t len) {
    (void)input;
    (void)len;
    return 0;
}

static long part2(const char *input, size_t len) {
    (void)input;
    (void)len;
    return 0;
}
