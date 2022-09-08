package main

type server interface {
	executeMain()
}

func run(s server) {
	s.executeMain()
}

func main() {
	bs := basicServcer{}
	run(bs)
}
