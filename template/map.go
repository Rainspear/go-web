package main

import (
	"os"
	"text/template"
)

type mapTemplate struct {
}

func (mapTemplate) executeMain() {
	m := map[string]string{"first": "Halo", "last": "Halo"}
	tpl := template.Must(template.ParseFiles("map.gohtml"))
	tpl.ExecuteTemplate(os.Stdout, "map.gohtml", m)
}
