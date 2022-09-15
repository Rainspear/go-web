package main

import (
	"bytes"
	"html/template"
	"io"
	"log"
	"net/http"
)

type queryState struct{}

type person struct {
	FirstName  string
	LastName   string
	Subscribed bool
}

func (queryState) executeMain() {
	http.Handle("/", http.HandlerFunc(home))
	http.Handle("/dog", http.HandlerFunc(dog))
	http.Handle("/favicon.ico", http.HandlerFunc(http.NotFound))
	log.Println(http.ListenAndServe(":8089", nil))
}

func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "text/html; charset=utf-8")
	f := r.FormValue("first")
	l := r.FormValue("last")
	s := r.FormValue("subscribe") == "on"
	// fmt.Println(r.FormValue("subscribe")) // will print "on"
	data := person{f, l, s}
	body := `
	<form method="POST">
    <label for="firstName">First Name</label>
    <input type="text" id="firstName" name="first">
    <br>
    <label for="lastName">Last Name</label>
    <input type="text" id="lastName" name="last">
    <br>
    <label for="sub">Subscribe</label>
    <input type="checkbox" id="sub" name="subscribe">
    <br>
    <input type="submit">
	</form>
	<br>
	<h1>First: {{.FirstName}}</h1>
	<h1>Last: {{.LastName}}</h1>
	<h1>Subscribed: {{.Subscribed}}</h1>
`
	tpl := template.Must(template.New("index.gohtml").Parse(body))
	tpl.ExecuteTemplate(w, "index.gohtml", data)
	// w.Write([]byte(body))
}

func dog(w http.ResponseWriter, r *http.Request) {
	r.ParseForm() // parse key value of all query and body request
	// v := r.FormValue("q") // parse key "q"
	var b bytes.Buffer
	for k, v := range r.Form {
		b.WriteString("Key is :" + k + " ;" + " Values are : ")
		for i, value := range v {
			if i < len(v)-1 {
				b.WriteString(value + ", ")
			} else {
				b.WriteString(value + ".")
			}
		}
		b.WriteString("\r \n")
	}
	io.WriteString(w, b.String())
}

// http:localhost:8089/?q=dog
