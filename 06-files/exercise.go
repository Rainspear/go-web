package main

import (
	"fmt"
	"log"
	"net/http"
)

type exerciseServer struct {
}

func foo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain;charset=utf-8")
	fmt.Fprintln(w, "Foo ran")
}

func execiseDog(w http.ResponseWriter, r *http.Request) {

}

func serveCorgi(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "/images/corgi.jpg")
}

func (exerciseServer) executeMain() {
	http.Handle("/", http.HandlerFunc(foo))
	http.HandleFunc("/dog/", execiseDog)
	log.Fatalln(http.ListenAndServe(":8089", nil))
}
