package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	"text/template"
)

type muxServer struct{}

type dataTemplate struct {
	Render         bool
	Method         string
	Title          string
	Para           string
	NestedTemplate string
}

func (muxServer) executeMain() {
	li, err := net.Listen("tcp", ":8089")
	if err != nil {
		log.Println(err)
	}
	fmt.Println("Server start on port 8089")
	for {
		conn, err := li.Accept()
		if err != nil {
			log.Println(err)
		}
		go handleMuxRequest(conn)
	}
}

func handleMuxRequest(conn net.Conn) {
	defer conn.Close()

	request(conn)

	// response(conn)
}

func request(conn net.Conn) {
	i := 0
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() { // loop every line of http request
		ln := scanner.Text()
		fmt.Println(ln)
		if i == 0 {
			mux(conn, ln)
		}
		if ln == "" {
			// headers are done
			break
		}
		i++
	}
}

// func response(conn net.Conn) {
// 	tpl := template.Must(template.New("mux").ParseFiles("mux.gohtml"))
// 	// fmt.Println(tpl.DefinedTemplates())
// 	// tpl.ExecuteTemplate(conn, "mux.gohtml", "Hello world") // write to body response with template
// 	fmt.Fprint(conn, "HTTP/1.1 200 OK\r\n")
// 	fmt.Fprintf(conn, "Content-Length: %d \r\n", 2048)
// 	fmt.Fprint(conn, "Content-Type: text/html \r\n")
// 	fmt.Fprint(conn, "\r\n")
// 	tpl.ExecuteTemplate(conn, "mux.gohtml", "Hello world 333") // write to body response with template
// }

func mux(conn net.Conn, line string) {
	m := strings.Fields(line)[0] // split by space " "
	u := strings.Fields(line)[1] // split by space " "
	data := dataTemplate{
		Render:         false,
		Method:         "GET",
		Title:          "Home",
		Para:           "",
		NestedTemplate: "",
	}
	// if m == "GET" && u == "/" {
	// 	renderHomePage(conn)
	// }
	if m == "GET" && u == "/about" {
		data.Title = "About"
		data.Render = true
		data.NestedTemplate = "About"
	}
	if m == "GET" && u == "/contact" {
		data.Title = "Contact"
		data.Render = true
		data.NestedTemplate = "Contact"
	}
	if m == "GET" && u == "/apply" {
		data.Title = "Apply"
		data.Render = true
		data.NestedTemplate = ""
	}
	if m == "POST" && u == "/apply" {
		data.Title = "Apply"
		data.Render = true
		data.NestedTemplate = "Apply"
	}
	renderPage(conn, data)
}

func renderPage(conn net.Conn, data dataTemplate) {
	tpl := template.Must(template.New("mux").ParseGlob("mux-*.gohtml"))
	fmt.Fprint(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length: %d \r\n", 2048)
	fmt.Fprint(conn, "Content-Type: text/html \r\n")
	fmt.Fprint(conn, "\r\n")
	tpl.ExecuteTemplate(conn, "mux.gohtml", data) // write to body response with template
}
