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
	defer li.Close()

	fmt.Println("Server start on port 8089")
	for {
		conn, err := li.Accept()
		if err != nil {
			log.Println(err.Error())
			continue
		}
		go handleMuxRequest(conn)
	}
}

func handleMuxRequest(conn net.Conn) {
	defer conn.Close()

	muxRequest(conn)

	// response(conn)
}

func muxRequest(conn net.Conn) {
	i := 0
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() { // loop every line of http request
		ln := scanner.Text()
		fmt.Println(ln)
		if i == 0 {
			mux(conn, ln)
			// m := strings.Fields(ln)[0]
			// u := strings.Fields(ln)[1]
			// fmt.Println("***METHOD", m)
			// fmt.Println("***URI", u)
		}
		if ln == "" {
			// headers are done
			break
		}
		i++
	}
}

// func response(conn net.Conn) {
// 	tpl := template.Must(template.New("mux-index.gohtml").ParseFiles("mux-index.gohtml"))
// 	// fmt.Println(tpl.DefinedTemplates())
// 	// tpl.ExecuteTemplate(conn, "mux.gohtml", "Hello world") // write to body response with template
// 	fmt.Fprint(conn, "HTTP/1.1 200 OK\r\n")
// 	fmt.Fprintf(conn, "Content-Length: %d \r\n", 2048)
// 	fmt.Fprint(conn, "Content-Type: text/html \r\n")
// 	fmt.Fprint(conn, "\r\n")
// 	tpl.ExecuteTemplate(conn, "mux-index.gohtml", nil) // write to body response with template
// }

func mux(conn net.Conn, line string) {
	m := strings.Fields(line)[0] // split by space " "
	u := strings.Fields(line)[1] // split by space " "
	// fmt.Println("Method", m)
	// fmt.Println("URI", u)
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
		data.NestedTemplate = "Apply"
	}
	if m == "POST" && u == "/apply" {
		data.Title = "Apply"
		data.Render = true
		data.NestedTemplate = ""
	}
	renderPage(conn, data)
}

func renderPage(conn net.Conn, data dataTemplate) {
	// tpl := template.Must(template.New("mux-index.gohtml").ParseGlob("mux-*.gohtml"))
	tplName := "mux-index.gohtml"
	tpl := template.Must(template.New(tplName).ParseGlob("mux-*.gohtml"))
	fmt.Fprint(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length: %d \r\n", 4096)
	fmt.Fprint(conn, "Content-Type: text/html \r\n")
	fmt.Fprint(conn, "\r\n")
	tpl.ExecuteTemplate(conn, tplName, data) // write to body response with template
	/*
		tplName should be the same when pass to "New" and ExecuteTemplate
		tplName is the name of index file and link to nested Template
	*/
}
