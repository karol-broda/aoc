# advent of code

my solutions to [advent of code](https://adventofcode.com/) puzzles written in go

## requirements 

this repo uses nix flakes and direnv for reproducible development environments

- [nix](https://nixos.org/download.html) with flakes enabled
- [direnv](https://direnv.net/)

## setup

clone the repo and allow direnv

the devshell includes:
- go (1.24.9)
- gopls (language server)
- gotools (goimports, etc)
- go-tools (staticcheck)

## running solutions

```bash
cd 2025/day01
go run .
```

or with example input:

```bash
go run . < example.txt
```

## license

solutions are my own. advent of code is created by [eric wastl](http://was.tl/)

