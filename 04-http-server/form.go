package main

import (
	"log"
	"net/http"
)

type formServer struct{}

type contactHandler struct{}

func (contactHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Fatalln(err)
	}
	tpl.ExecuteTemplate(w, "index.gohtml", r.Form)
}

func (formServer) executeMain() {
	ch := contactHandler{}
	http.ListenAndServe(":8089", ch)
}
