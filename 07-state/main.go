package main

type stateServer interface {
	executeMain()
}

func run(s stateServer) {
	s.executeMain()
}

func main() {
	// qs := queryState{}
	// run(qs)

	// fs := fileState{}
	// run(fs)

	cs := cookieState{}
	run(cs)
}
