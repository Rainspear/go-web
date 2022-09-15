package main

import (
	"bytes"
	"io"
	"log"
	"net/http"
)

type queryState struct{}

func (queryState) executeMain() {
	http.Handle("/", http.HandlerFunc(dog))
	http.Handle("/favicon.ico", http.HandlerFunc(http.NotFound))
	log.Println(http.ListenAndServe(":8089", nil))
}

func dog(w http.ResponseWriter, r *http.Request) {
	r.ParseForm() // parse key value of all query and body request
	// v := r.FormValue("q") // parse key "q"
	var b bytes.Buffer
	for k, v := range r.Form {
		b.WriteString("Key is :" + k + " ;" + " Values are : ")
		for i, value := range v {
			if i < len(v)-1 {
				b.WriteString(value + ", ")
			} else {
				b.WriteString(value + ".")
			}
		}
		b.WriteString("\r \n")
	}
	io.WriteString(w, b.String())
}

// http:localhost:8089/?q=dog
