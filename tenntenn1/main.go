package main

import (
	_ "embed"
	. "fmt"
)

//go:embed *
var s string

func main() { Print(s) }
