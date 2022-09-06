package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
)

type handleServer struct {
}

func (handleServer) executeMain() {
	li, err := net.Listen("tcp", ":8089")
	if err != nil {
		log.Println(err)
	}

	fmt.Println("Server starts on port 8089")

	for {
		con, err := li.Accept()
		if err != nil && 1 == 2 {
			log.Println(err)
		}
		go handleRequest(con)
	}
}

func handleRequest(conn net.Conn) {
	defer conn.Close()

	request(conn)

	response(conn)
}

func request(conn net.Conn) {
	i := 0
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		ln := scanner.Text()
		fmt.Println(ln)
		// fmt.Println("field: ", strings.Fields(ln))
		if i == 0 {
			// request line
			m := strings.Fields(ln)[0]
			u := strings.Fields(ln)[1]
			fmt.Println("***METHOD", m)
			fmt.Println("***URI", u)
		}
		if ln == "" {
			// headers are done
			break
		}
		i++
	}
}

func response(conn net.Conn) {
	body := `<!DOCTYPE html>
	<html>
		<head>
			<meta charset="utf-8">
			<meta http-equiv="X-UA-Compatible" content="IE=edge">
			<title></title>
			<meta name="description" content="">
			<meta name="viewport" content="width=device-width, initial-scale=1">
			<link rel="stylesheet" href="">
		</head>
		<body>
			<strong>Hello world 123213213</strong>
		</body>
	</html>
	`

	fmt.Fprint(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length: %d \r\n", len(body))
	fmt.Fprint(conn, "Content-Type: text/html \r\n")
	fmt.Fprint(conn, "\r\n")
	fmt.Fprint(conn, body)

	// bs := make([]byte, 2048)
	// bs = append(bs, `<!DOCTYPE html>
	// <html>
	// 	<head>
	// 		<meta charset="utf-8">
	// 		<meta http-equiv="X-UA-Compatible" content="IE=edge">
	// 		<title></title>
	// 		<meta name="description" content="">
	// 		<meta name="viewport" content="width=device-width, initial-scale=1">
	// 		<link rel="stylesheet" href="">
	// 	</head>
	// 	<body>
	// 		<strong>Hello world</strong>
	// 	</body>
	// </html>
	// `...)
	// nbs := make([]byte, 2048)
	// nbs = append(nbs, "3213213"...)
	// conn.Write(nbs)
	// conn.Read(nbs)
	// fmt.Println("Body response: ", string(nbs))
	// bs := make([]byte, 2048)
	// conn.Read(bs)
	// fmt.Printf(string(bs))
}
