package main

type session interface {
	executeMain()
}

func run(s session) {
	s.executeMain()
}

func main() {
	// is := uuidSession{}
	// run(is)

	us := userSession{}
	run(us)
}
