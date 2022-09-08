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
	tpl.ExecuteTemplate(w, "index.gohtml", r.Form) // data in body and query
	// tpl.ExecuteTemplate(w, "index.gohtml", r.PostForm) // data only in body

}

func (formServer) executeMain() {
	ch := contactHandler{}
	err := http.ListenAndServe(":8089", ch)
	if err != nil {
		log.Println(err.Error())
	}
}
