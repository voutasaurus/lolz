package main

import (
	"fmt"
	"text/template"
	"os"
)

func main() {
	str := "package main\n\nimport (\n\t\"fmt\"\n\t\"text/template\"\n\t\"os\"\n)\n\nfunc main() {\n\tstr := {{.}}\n\ttmpl, _ := template.New(\"test\").Parse(str)\n\t_ = tmpl.Execute(os.Stdout, fmt.Sprintf(\"%q\", str))\n}"
	tmpl, _ := template.New("test").Parse(str)
	_ = tmpl.Execute(os.Stdout, fmt.Sprintf("%q", str))
}