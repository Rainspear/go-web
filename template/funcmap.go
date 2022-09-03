package main

import (
	"os"
	"strings"
	"text/template"
)

type funcMapTemplate struct {
}

func firstThreeString(s string) string {
	ss := strings.TrimSpace(s)
	return ss[:3]
}

func (funcMapTemplate) executeMain() {
	s := "This is original string"
	// this code below is not working because the funcmap is called after parsing file
	// => so when parsing file there is no funcmap in the template
	// => cause undifined funcmap (ft uc function)
	// tpl := template.Must(template.ParseFiles("./funcmap.gohtml"))
	// tpl.Funcs(template.FuncMap{
	// 	"ft": firstThreeString,
	// 	"up": strings.ToUpper,
	// })
	tpl := template.Must(template.New("").Funcs(template.FuncMap{
		"ft": firstThreeString,
		"up": strings.ToUpper,
	}).ParseFiles("funcmap.gohtml"))
	tpl.ExecuteTemplate(os.Stdout, "funcmap.gohtml", s)

}
