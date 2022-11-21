package main

import . "fmt"

var s = "package main\n\nimport . \"fmt\"\n\nvar s = %q\n\nfunc main() { Printf(s, s) }\n"

func main() { Printf(s, s) }
