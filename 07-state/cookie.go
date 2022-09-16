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
	http.Handle("/expire", http.HandlerFunc(expireCookie))
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8089", nil)
}

func expireCookie(w http.ResponseWriter, r *http.Request) {
	for _, v := range r.Cookies() {
		v.MaxAge = -1 // < 0 will delete the cookie
		http.SetCookie(w, v)
	}
	fmt.Fprintf(w, "Cookies are expired\n")
	http.Redirect(w, r, "/read", http.StatusSeeOther)
}

func readCookie(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Cookie: %+v\n", r.Cookies())
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
