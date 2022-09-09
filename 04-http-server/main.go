package main

import (
	"html/template"
)

type server interface {
	executeMain()
}

func run(s server) {
	s.executeMain()
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("index.gohtml"))
}

func main() {
	// bs := basicServcer{}
	// run(bs)

	fs := formServer{}
	run(fs)
}
