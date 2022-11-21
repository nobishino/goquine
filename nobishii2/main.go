package main

import f "fmt"

func main() { f.Printf("%s%q\n", q, q) }

var q = "package main\n\nimport f \"fmt\"\n\nfunc main() { f.Printf(\"%s%q\\n\", q, q) }\n\nvar q = "
