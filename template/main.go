package main

type commonTemplate interface {
	executeMain()
}

func run(c commonTemplate) {
	c.executeMain()
}

func main() {
	vt := variableTemplate{}
	run(vt)

	rt := rangeTemplate{}
	run(rt)

	mt := mapTemplate{}
	run(mt)

	st := structTemplate{}
	run(st)

	ft := funcMapTemplate{}
	run(ft)

	pt := pipelineTemplate{}
	run(pt)
}
