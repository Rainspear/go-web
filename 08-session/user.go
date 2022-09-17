package main

import (
	"net/http"
	"text/template"
)

type userSession struct{}

func (userSession) executeMain() {
	http.Handle("/", http.HandlerFunc(home))
	http.Handle("/login", http.HandlerFunc(login))
}

type User struct {
	FullName string
	Password string
	UserName string
}

func login(w http.ResponseWriter, r *http.Request) {
	var data User
	if r.Method == http.MethodPost {
		err := r.ParseForm()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		userName := r.FormValue("username")
		password := r.FormValue("password")
		fullName := r.FormValue("fullname")

		data := User{
			userName,
			password,
			fullName,
		}

	}
	tpl := template.Must(template.New("index.go").ParseFiles("./template/index.gohtml"))
	tpl.ExecuteTemplate(w, "index.go", data)
}

func home(w http.ResponseWriter, r *http.Request) {
	// id := uuid.New()

}
