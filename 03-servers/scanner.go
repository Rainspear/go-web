package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

type scannerServer struct{}

func (scannerServer) executeMain() {
	li, err := net.Listen("tcp", ":8089")
	if err != nil {
		log.Panic(err)
	}
	defer li.Close()

	for {
		con, err := li.Accept()
		if err != nil {
			log.Fatalln(err)
		}
		go handle(con)
	}
}

func handle(conn net.Conn) {
	scanner := bufio.NewScanner(conn)

	for scanner.Scan() {
		ln := (scanner.Text())
		fmt.Println(ln)
	}
	defer conn.Close()

	// never reach here
}
