package main

import . "fmt"

func main() {
	q := `package main

import . "fmt"

func main() {
	q := %c%s%[1]c
	Printf(q, 96, q)
}
`
	Printf(q, 96, q)
}
