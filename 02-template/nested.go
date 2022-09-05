package main

import (
	"os"
	"text/template"
)

type nestedTemplate struct{}

type personTest struct {
	Name string
	Age  int
}

func (p personTest) TakeArg(x int) int {
	return 2 * x
}

func (p personTest) AgeDb() int {
	return p.Age * 2
}

func (nestedTemplate) executeMain() {
	p := personTest{
		Name: "John",
		Age:  32,
	}
	tpl := template.Must(template.New("nested.gohtml").ParseGlob("nested*.gohtml"))
	tpl.ExecuteTemplate(os.Stdout, "nested.gohtml", p)
}
