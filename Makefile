CC = zig cc
CFLAGS = -Wall -Wextra -O2 -std=c11

YEAR ?= 2025
DAY ?= 01

SRC = $(YEAR)/day$(DAY)/main.c
BIN = $(YEAR)/day$(DAY)/main

.PHONY: build run clean

build:
	@$(CC) $(CFLAGS) -o $(BIN) $(SRC)

run: build
	@cd $(YEAR)/day$(DAY) && ./main

clean:
	@find . -name "main" -type f -executable -delete

