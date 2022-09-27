package main

import (
	"net/http"
	"strings"
	"text/template"

	"github.com/google/uuid"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {
	http.Handle("/", http.HandlerFunc(index))
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8089", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	c := getCookie(w, r)
	c = appendValue(w, c)
	xs := strings.Split(c.Value, "|")
	tpl.ExecuteTemplate(w, "index.gohtml", xs)
}

func getCookie(w http.ResponseWriter, r *http.Request) *http.Cookie {
	c, err := r.Cookie("session")
	if err != nil {
		id := uuid.New().String()
		c = &http.Cookie{
			Name:  "session",
			Value: id,
		}
		http.SetCookie(w, c)
	}
	return c
}

func appendValue(w http.ResponseWriter, c *http.Cookie) *http.Cookie {
	p1 := "disneyland.jpg"
	p2 := "atbeach.jpg"
	p3 := "hollywood.jpg"
	s := c.Value
	if !strings.Contains(s, p1) {
		s += "|" + p1
	}
	if !strings.Contains(s, p2) {
		s += "|" + p2
	}
	if !strings.Contains(s, p3) {
		s += "|" + p3
	}
	c.Value = s
	http.SetCookie(w, c)
	return c
}
