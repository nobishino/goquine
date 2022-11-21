package main

import . "fmt"

func main() {
	a := "package main\n\nimport . \"fmt\"\n\nfunc main(){\n\ta:=%q\n\tPrintf(a, a)\n}"
	Printf(a, a)
}

