package main

type commonServer interface {
	executeMain()
}

func run(c commonServer) {
	c.executeMain()
}

func main() {
	// ts := telnetServer{}
	// run(ts)

	// ss := scannerServer{}
	// run(ss)

	// hs := handleServer{}
	// run(hs)

	ms := muxServer{}
	run(ms)
}
