package main

import (
	"fmt"
	"net/http"
	"text/template"

	"github.com/google/uuid"
)

type userSession struct{}

var tpl *template.Template
var dbUsers = map[string]User{}      // user ID, user
var dbSessions = map[string]string{} // session ID, user ID

func init() {
	tpl = template.Must(template.ParseGlob("./templates/*"))
}

func (userSession) executeMain() {
	http.Handle("/login", http.HandlerFunc(login))
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.Handle("/", http.HandlerFunc(index)) // remember to add this line under other routes
	http.ListenAndServe(":8089", nil)
}

type User struct {
	UserName string
	Password string
	First    string
	Last     string
}

func login(w http.ResponseWriter, r *http.Request) {
	var data User

	// extract session from cookie
	c, err := r.Cookie("session")
	if err != nil { // if can not parse cookie (empty cookie with key) -> create new session
		sid := uuid.New()
		c = &http.Cookie{
			Name:     "session",
			Value:    sid.String(),
			HttpOnly: true, // for not modifying the session by javscript
			// Secure: true,
		}
		http.SetCookie(w, c) // add session to cookie -> response: Set-Cookie: session; expired=50ms;
	}

	// extract data user from session
	if userId, ok := dbSessions[c.Value]; ok {
		data = dbUsers[userId]
	}

	if r.Method == http.MethodPost {
		// parse form data
		err := r.ParseForm()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		u := r.FormValue("username")
		p := r.FormValue("password")
		f := r.FormValue("firstname")
		l := r.FormValue("lastname")
		data = User{u, p, f, l}

		// save new userId to session
		dbSessions[c.Value] = u

		// save data user
		dbUsers[u] = data
	}

	// return template
	// tpl := template.Must(template.New("login.gohtml").ParseFiles("./template/login.gohtml"))
	tpl.ExecuteTemplate(w, "login.gohtml", data)
}

func index(w http.ResponseWriter, r *http.Request) {
	var data User
	// extract session from cookie
	c, err := r.Cookie("session")
	if err != nil { // if can not parse cookie (not authenicated) -> redirect to /login
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	// extract user data from session
	uid, ok := dbSessions[c.Value]
	if !ok { // can not get user from session
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}
	data = dbUsers[uid]
	fmt.Printf("data: %+v\n", data)

	// return template
	// tpl := template.Must(template.New("index.gohtml").ParseFiles("./template/index.gohtml"))
	tpl.ExecuteTemplate(w, "index.gohtml", data)
}
