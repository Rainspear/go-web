package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

type fileState struct{}

func (fileState) executeMain() {
	http.Handle("/", http.HandlerFunc(cat))
	http.ListenAndServe(":8089", nil)
}

func cat(w http.ResponseWriter, r *http.Request) {
	var s string
	if r.Method == http.MethodPost {
		// open
		f, h, err := r.FormFile("q")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer f.Close()

		fmt.Println("\nfile:", f, "\nheader:", h, "\nerr", err)

		// read
		bs, err := ioutil.ReadAll(f)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		s = string(bs)
	}
	w.Header().Add("Content-Type", "text/html; charset=utf-8;")
	io.WriteString(w, `
	<form method="POST" enctype="multipart/form-data">
	<input type="file" name="q">
	<input type="submit">
	</form>
	<br>`+s)
}
