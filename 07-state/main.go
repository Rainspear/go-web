package main

type stateServer interface {
	executeMain()
}

func run(s stateServer) {
	s.executeMain()
}

func main() {
	qs := queryState{}
	run(qs)
}
