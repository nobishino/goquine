package main

import "fmt"

func main() {
	fmt.Println(q + fmt.Sprintf("%q", q))
}

var q = "package main\n\nimport \"fmt\"\n\nfunc main() {\n\tfmt.Println(q + fmt.Sprintf(\"%q\", q))\n}\n\nvar q = "
