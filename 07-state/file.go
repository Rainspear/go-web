package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

type fileState struct{}

func (fileState) executeMain() {
	http.Handle("/", http.HandlerFunc(cat))
	http.Handle("/image", http.HandlerFunc(serveDogImage))
	http.Handle("/static", http.StripPrefix("/static", http.FileServer(http.Dir("./assets"))))
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

func serveDogImage(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		f, h, err := r.FormFile("q")
		if err != nil {
			// return error response to user
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		fmt.Println("\nfile:", f, "\nheader:", h, "\nerr", err)
	}
	// create file
	// t := strings.Join(strings.Split(time.Now().Format(time.Stamp), " "), "-") //  Nov-10-23:00:00
	t := time.Now().Format(time.RFC3339) // 2009-11-10T23:00:00Z
	fmt.Println("time", t)
	name := string([]byte(`dog` + t + ".jpg"))
	os.WriteFile(name, f, 0666)
	// nf, err := os.Create(string(name))
	// if err != nil {
	// 	http.Error(w, err.Error(), http.StatusInternalServerError)
	// 	log.Println("Error when create file")
	// }

	// response
	w.Header().Add("Content-Type", "text/html; charset=utf-8;")
	io.WriteString(w, `
	<form method="POST" enctype="multipart/form-data">
	<input type="file" name="q">
	<input type="submit">
	</form>

	<br>`+s)
}
