package main

import (
	"os"
	"text/template"
)

type person struct {
	Name string
	Age  int
	City string
}

type structTemplate struct {
}

func (structTemplate) executeMain() {
	p := person{Name: "John", Age: 1, City: "London"}
	tpl := template.Must(template.ParseFiles("struct.gohtml"))
	tpl.ExecuteTemplate(os.Stdout, "struct.gohtml", p)
}
