![Test](https://github.com/nobishino/goquine/actions/workflows/test.yml/badge.svg)
# Goquine

Goquine is a repository to share [Quine](https://en.wikipedia.org/wiki/Quine_(computing)) programs written in Go.

## Rule

- Output to `stdout` (not `stderr`)
- No 3rd party library

## Contribute

- create directory `<your name>+positive integer`
    - example: `nobishii1`
- create `main.go` and write your quine code
- run `make` to ensure your code is go formatted and is quine
- make a pull request

## Current records

### Shortest quine (with standard library)

[tenntenn1](./tenntenn1/main.go) (110 bytes)

### Shortest quine (with standard library, without embed)

[nobishii2](./nobishii2/main.go) (174 bytes)


### Shortest quine (without standard library)

NONE