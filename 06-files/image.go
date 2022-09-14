package main

import (
	"io"
	"log"
	"net/http"
	"os"
)

type imageServer struct{}

func d(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html;charset=utf-8")
	io.WriteString(w, `
		<img src='https://cdn.pixabay.com/photo/2019/08/19/07/45/corgi-4415649_960_720.jpg' alt="corgi_dog" />
	`)
}

func dogPig(w http.ResponseWriter, r *http.Request) {
	f, err := os.Open("/images/corgi.jpg")
	if err != nil {
		http.Error(w, "file not found", 404)
		return
	}
	defer f.Close()

	io.Copy(w, f)
}

func serveDog(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "/images/corgi.jpg")
}

func serveContent(w http.ResponseWriter, r *http.Request) {
	f, err := os.Open("/images/corgi.jpg")
	if err != nil {
		http.Error(w, "file not found", 404)
		return
	}
	defer f.Close()

	fi, err := f.Stat()
	if err != nil {
		http.Error(w, "file not found", 404)
		return
	}

	http.ServeContent(w, r, fi.Name(), fi.ModTime(), f)
}

func (imageServer) executeMain() {
	http.Handle("/", http.FileServer(http.Dir(".")))
	http.Handle("/static/", http.StripPrefix("/static", http.FileServer(http.Dir("./images")))) // serve request as static
	http.Handle("/test", http.HandlerFunc(d))                                                   // route must be "/"
	http.Handle("/dog", http.HandlerFunc(dogPig))
	http.Handle("/serve-dog", http.HandlerFunc(serveDog))
	http.Handle("/serve-content", http.HandlerFunc(serveContent))
	err := http.ListenAndServe(":8089", nil)
	if err != nil {
		log.Println("Can not start server on port 8089")
	}
}
