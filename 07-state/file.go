package main

import "net/http"

type fileState struct{}

func (fileState) executeMain() {
	http.Handle("/", http.HandlerFunc(cat))
	http.ListenAndServe(":8089", nil)
}

func cat(w http.ResponseWriter, r *http.Request) {

}
