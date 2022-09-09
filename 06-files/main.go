package main

type fileServer interface {
	executeMain()
}

func run(f fileServer) {
	f.executeMain()
}

func main() {
	is := imageServer{}
	run(is)
}
