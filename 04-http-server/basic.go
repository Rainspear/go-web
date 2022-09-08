package main

import (
	"fmt"
	"log"
	"net/http"
)

type basicServcer struct{}

type contactHandler struct{}

func (contactHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Server starts on port 8089")
	_, err := w.Write([]byte(`
	<!DOCTYPE html>
	<html>
		<head>
			<meta charset="utf-8">
			<meta http-equiv="X-UA-Compatible" content="IE=edge">
			<title></title>
			<meta name="description" content="">
			<meta name="viewport" content="width=device-width, initial-scale=1">
			<link rel="stylesheet" href="">
		</head>
		<body>
				<a href="/">index</a><br>
				<a href="/about">about</a><br>
				<a href="/contact">contact</a><br>
				<a href="/apply">apply</a><br>
			<h1> Mux Template World! </h1>
			<script src="" async defer></script>
		</body>
	</html>
	`))
	if err != nil {
		log.Println(err, "Can not write to response")
	}
}

func (basicServcer) executeMain() {
	c := contactHandler{}
	err := http.ListenAndServe(":8089", c)
	if err != nil {
		log.Println(err.Error())
	}
}
