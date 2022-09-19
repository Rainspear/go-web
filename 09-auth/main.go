package main

import (
	"net/http"
	"text/template"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type user struct {
	UserName string
	Password []byte
	First    string
	Last     string
}

type session struct {
	userId       string
	lastActivity time.Time
}

var tpl *template.Template
var dbUsers map[string]user
var dbSessions map[string]session

const sessionLength int = 30 // 30 seconds

func init() {
	tpl = template.Must(template.ParseGlob("./templates/*"))
}

func main() {
	http.Handle("/signin", http.HandlerFunc(signin))
	http.Handle("/signup", http.HandlerFunc(signup))
	http.Handle("/authorize", http.HandlerFunc(authorize))
	http.Handle("/", http.HandlerFunc(index))
	http.ListenAndServe(":8089", nil)
}

func index(w http.ResponseWriter, r *http.Request) {

	tpl.ExecuteTemplate(w, "index.gohtml", nil)
}

func signin(w http.ResponseWriter, r *http.Request) {
	var data user
	if r.Method == http.MethodPost {
		u := r.FormValue("username")
		p := r.FormValue("password")
		bs := []byte(p)
		bcrypt.GenerateFromPassword(bs, bcrypt.MinCost)
		data, ok := dbUsers[u]
		if !ok {
			http.Error(w, "username or password is not correct", http.StatusBadRequest)
			return
		}
		err := bcrypt.CompareHashAndPassword(data.Password, bs)
		if err != nil {
			http.Error(w, "username or password is not correct", http.StatusBadRequest)
			return
		}
		sid := uuid.New().String()
		c := &http.Cookie{
			Name:   "session",
			Value:  sid,
			MaxAge: sessionLength,
		}
		http.SetCookie(w, c)
		dbSessions[c.Value] = session{u, time.Now()}
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	tpl.ExecuteTemplate(w, "signin.gohtml", data)
}

func signup(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "signup.gohtml", nil)
}

func authorize(w http.ResponseWriter, r *http.Request) {

}
