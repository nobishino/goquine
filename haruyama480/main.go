package main

import (
	. "fmt"
	. "os/exec"
)

func main() { c, _ := Command("cat", "haruyama480/main.go").Output(); Print(string(c)) }
