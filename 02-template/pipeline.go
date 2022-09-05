package main

import (
	"math"
	"os"
	"text/template"
)

type pipelineTemplate struct {
}

func square(x int) int { return x * x }

func double(x int) int { return 2 * x }

func sqrt(x int) int { return int(math.Sqrt(float64(x))) }

var fm = template.FuncMap{
	"sq":   square,
	"db":   double,
	"sqrt": sqrt,
}

func (pipelineTemplate) executeMain() {
	n := 100
	tpl := template.Must(template.New("").Funcs(fm).ParseFiles("pipeline.gohtml"))
	tpl.ExecuteTemplate(os.Stdout, "pipeline.gohtml", n)
}
