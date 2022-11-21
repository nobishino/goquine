![Test](https://github.com/nobishino/goquine/actions/workflows/test.yml/badge.svg)
# Goquine

Goquine is a repository to share [Quine](https://en.wikipedia.org/wiki/Quine_(computing)) programs written in Go.

## Regulations

- Output to `stdout`. (not `stderr`)
- No 3rd party library.
- No `os.Open` or such external input.

If you want to submit program which does not conform to these regulations, please submit to `unofficial/<yourname> + positive integer/main.go`.

## Contribution

- create directory `<your name>+positive integer`
    - example: `nobishii1`
- create `main.go` and write your quine code
- run `make` to ensure your code is go formatted and is quine
- make a pull request

## Current records

```
find ./*/main.go -type f | xargs -I{} wc -c {}
     150 ./cia-rana1/main.go
     200 ./nobishii1/main.go
     174 ./nobishii2/main.go
     108 ./tenntenn1/main.go
     151 ./tenntenn2/main.go
find ./unofficial/*/main.go -type f | xargs -I{} wc -c {}
     126 ./unofficial/tenntenn3/main.go
find ./unofficial/stderr/*/main.go -type f | xargs -I{} wc -c {}
     137 ./unofficial/stderr/cia-rana2/main.go
```

### Shortest quine (with standard library)

- [tenntenn1](./tenntenn1/main.go) (108 bytes)

### Shortest quine (with standard library, without embed)

- [cia-rana1](./cia-rana1/main.go) (150 bytes)

### Shortest quine (without standard library)

NONE

### Unofficial records(not conforming to the regulations)

- [unofficial/tenntenn3](./unofficial/tenntenn3/main.go) (126 bytes, using external input)
- [unofficial/stderr/cia-rana2](./unofficial/stderr/cia-rana2/main.go) (137 bytes, using stderr instead of stdout)