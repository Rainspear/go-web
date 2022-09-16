package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

type cookieState struct{}

func (cookieState) executeMain() {
	http.Handle("/set", http.HandlerFunc(setCookie))
	http.Handle("/read", http.HandlerFunc(readCookie))
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8089", nil)
}

func readCookie(w http.ResponseWriter, r *http.Request) {
	c := r.Cookies()
	fmt.Fprintf(w, "Cookie: %+v\n", c)
}

func setCookie(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Printf("Error parsing form: %s\n", err.Error())
	}
	for k, v := range r.Form {
		fmt.Printf("key: %s - values: %+v\n", k, v)
		http.SetCookie(w, &http.Cookie{
			Name:  k,
			Value: strings.Join(v, "-"),
		})
	}
	c := r.Cookies()
	if len(c) == 0 {
		http.SetCookie(w, &http.Cookie{
			Name:  "name",
			Value: "Kanz",
		})
	}
	fmt.Fprintf(w, "HTTP/1.1 Cookie is in browser now")
}
