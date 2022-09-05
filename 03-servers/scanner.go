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
	// err := conn.SetDeadline(time.Now().Add(10 * time.Second))
	// if err != nil {
	// 	log.Println("Conn TIMEOUT")
	// }
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		ln := (scanner.Text())
		// io.WriteString(conn, "Hello")
		fmt.Println(ln)
	}
	defer conn.Close()

	// never reach here
}
