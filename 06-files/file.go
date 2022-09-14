package main

import (
	"log"
	"net/http"
)

type fileContentServer struct{}

func (fileContentServer) executeMain() {
	err := http.ListenAndServe(":8089", http.FileServer(http.Dir(".")))
	if err != nil {
		log.Println(err)
	}
}
