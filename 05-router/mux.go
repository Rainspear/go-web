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

func (muxServer) executeMain() {
	var c contact
	var a about
	mux := http.NewServeMux() // it is still Handler type
	mux.Handle("/contact", c)
	mux.Handle("/about", a)

	err := http.ListenAndServe(":8089", mux)
	if err != nil {
		log.Println(err.Error())
	}
}
