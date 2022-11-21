package main

import (
	. "os"
	. "runtime"
)

func main() {
	_, f, _, _ := Caller(0)
	s, _ := ReadFile(f)
	Stdout.Write(s)
}
