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

// Can implicity understand as below code
// So that handlerfunc can be taken as argument of handlefunc
// right now handlerfunc is still handler type because it implements ServerHTTP method
// func handle(type handler interface -> ServeHTTP(w, r))
// func handlefunc(type handlerfunc func -> ServeHTTP(w, r))
/*
type handlerfunc func(w http.ResponseWriter, r *http.Request)
func (h handlerfunc) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h(w,r)
}
*/

func (contact) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Contact Handler")
	io.WriteString(w, "Contact Handler")
}

func (about) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println("About Handler")
	io.WriteString(w, "About Handler")
}

func contactHandler(w http.ResponseWriter, r *http.Request) { // this is called handlerfunc.
	fmt.Println("About Handler with handlefunc mux and default")
	io.WriteString(w, "About Handler handlefunc mux and default")
}

func aboutHandler(w http.ResponseWriter, r *http.Request) { // this is called handlerfunc
	fmt.Println("About Handler with handlefunc mux and default")
	io.WriteString(w, "About Handler handlefunc mux and default")
}

func (muxServer) executeMain() {
	// new mux
	// var c contact
	// var a about
	// mux := http.NewServeMux()  // it is still Handler type
	// mux.Handle("/contact/", c) // match relative path
	// mux.Handle("/about", a)    // match exact path
	// err := http.ListenAndServe(":8089", mux)

	// new mux with handlefunc
	mux := http.NewServeMux()
	mux.HandleFunc("/contact/", contactHandler) // this is handle func that take HandlerFunc as argument
	mux.HandleFunc("/about", aboutHandler)      // this is handle func that take HandlerFunc as argument
	err := http.ListenAndServe(":8089", mux)

	// default mux
	// http.HandleFunc("/contact/", contactHandler)
	// http.HandleFunc("/about", aboutHandler)
	// err := http.ListenAndServe(":8089", nil)

	// common code
	if err != nil {
		log.Println(err.Error())
	}
}
