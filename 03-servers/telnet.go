package main

import (
	"fmt"
	"io"
	"log"
	"net"
)

type telnetServer struct{}

func (telnetServer) executeMain() {
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
		io.WriteString(con, "Connected to server\n")
		fmt.Println("Server starts on port 8089")
		fmt.Println("Say hi")
		fmt.Printf("%v", "Hi There")
		con.Close() // temporarily invoke
	}
}
