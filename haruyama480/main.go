package main

import (
	. "fmt"
	. "os/exec"
)

func main() { c, _ := Command("cat", "main.go").Output(); Print(string(c)) }
