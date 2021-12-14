package main

import (
	"os"
	"text/template"
)

type Student struct {
	ID      uint
	Name    string
}

func main() {
	stu := Student{0, "jason"}
	tmpl, err := template.New("test").Parse("The name for student {{.ID}} is {{.Name}}")
	if err != nil { panic(err) }
	err = tmpl.Execute(os.Stdout, stu)
	if err != nil { panic(err) }
}
