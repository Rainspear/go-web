package main

import (
	"os"
	"text/template"
)

type rangeTemplate struct {
}

func (rangeTemplate) executeMain() {
	s := []string{"1", "2", "3", "4", "5", "6"}
	tpl := template.Must(template.ParseFiles("range.gohtml"))
	tpl.ExecuteTemplate(os.Stdout, "range.gohtml", s)
}
