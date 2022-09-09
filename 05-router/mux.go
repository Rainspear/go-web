package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

type muxServer struct{}

type contact int
type about int

func (contact) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Contact Handler")
	io.WriteString(w, "Contact Handler")
}

func (about) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println("About Handler")
	io.WriteString(w, "About Handler")
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("About Handler with default")
	io.WriteString(w, "About Handler with default")
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("About Handler with default")
	io.WriteString(w, "About Handler with default")
}

func (muxServer) executeMain() {
	// new mux
	// var c contact
	// var a about
	// mux := http.NewServeMux()  // it is still Handler type
	// mux.Handle("/contact/", c) // match relative path
	// mux.Handle("/about", a)    // match exact path
	// err := http.ListenAndServe(":8089", mux)

	// default mux
	http.HandleFunc("/contact/", contactHandler)
	http.HandleFunc("/about", aboutHandler)
	err := http.ListenAndServe(":8089", nil)
	if err != nil {
		log.Println(err.Error())
	}
}
