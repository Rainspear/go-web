package main

import (
	"os"
	"text/template"
)

type variableTemplate struct {
}

func (variableTemplate) executeMain() {
	// var t = template.Must(template.New("name").Parse("text")) // create new memory for template then parse "text" to it
	tpl := template.Must(template.ParseFiles("variable.gohtml"))
	tpl.ExecuteTemplate(os.Stdout, "variable.gohtml", "hello world")
}
