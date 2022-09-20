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
	Role     string
}

type session struct {
	userId       string
	lastActivity time.Time
}

var tpl *template.Template
var dbUsers = map[string]user{}
var dbSessions = map[string]session{}
var dbSessionsCleaned time.Time

const sessionLength int = 30 // 30 seconds

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
	btest, _ := bcrypt.GenerateFromPassword([]byte("123"), bcrypt.MinCost)
	dbUsers["kanz"] = user{"kanz", btest, "Kanz", "Han", "007"}
	dbSessionsCleaned = time.Now() // in real production it should clear when not in busy hour
}

func main() {
	http.Handle("/bar", http.HandlerFunc(bar))
	http.Handle("/signin", http.HandlerFunc(signin))
	http.Handle("/signup", http.HandlerFunc(signup))
	http.Handle("/signout", authorized(signout))
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.Handle("/", http.HandlerFunc(index))
	http.ListenAndServe(":8089", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	data := getUserData(w, r)
	showSessions()
	tpl.ExecuteTemplate(w, "index.gohtml", data)
}

func signin(w http.ResponseWriter, r *http.Request) {
	if alreadyLoggedIn(w, r) {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	var data user
	if r.Method == http.MethodPost {
		u := r.FormValue("username")
		p := r.FormValue("password")
		data, ok := dbUsers[u]
		if !ok {
			http.Error(w, "username or password is not correct", http.StatusBadRequest)
			return
		}
		err := bcrypt.CompareHashAndPassword(data.Password, []byte(p))
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

func signup(w http.ResponseWriter, req *http.Request) {
	var data user
	if alreadyLoggedIn(w, req) {
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}
	if req.Method == http.MethodPost {
		u := req.FormValue("username")
		p := req.FormValue("password")
		f := req.FormValue("firstname")
		l := req.FormValue("lastname")
		r := req.FormValue("role")

		if _, ok := dbUsers[u]; ok {
			http.Error(w, "Username already taken", http.StatusForbidden)
			return
		}
		b, err := bcrypt.GenerateFromPassword([]byte(p), bcrypt.MinCost)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		sid := uuid.New().String()
		c := &http.Cookie{
			Name:   "session",
			Value:  sid,
			MaxAge: sessionLength,
		}
		data := user{u, b, f, l, r}
		dbSessions[c.Value] = session{u, time.Now()}
		dbUsers[u] = data
		http.SetCookie(w, c)
		http.Redirect(w, req, "/", http.StatusSeeOther)
	}
	tpl.ExecuteTemplate(w, "signup.gohtml", data)
}

func signout(w http.ResponseWriter, r *http.Request) {
	c, _ := r.Cookie("session")
	delete(dbSessions, c.Value)
	c = &http.Cookie{Name: "", Value: "", HttpOnly: true, MaxAge: -1}
	http.SetCookie(w, c)

	// this cleanup function should run as cron job in background in a certain unbusy time in a day
	if time.Now().Sub(dbSessionsCleaned) > (time.Second * 30) { // clear after 30s from start and whatever people logout
		go cleanSessions() //
	}

	http.Redirect(w, r, "/signin", http.StatusSeeOther)
}

// middleware authorized for all routes, if not redirect to login
func authorized(h http.HandlerFunc) http.HandlerFunc { // middleware can apply for all route
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !alreadyLoggedIn(w, r) {
			http.Redirect(w, r, "/signin", http.StatusSeeOther)
			return
		}
		/*
			because definition of http.HandlerFunc is: type HandlerFunc func(w http.ResponseWritter, r *http.Request)
			and HandlerFunc override ServeHTTP(w,r) => Handler interface
			h = signout
		*/
		h.ServeHTTP(w, r)
	})
}

// route for authorized function, only "007" role can access data
func bar(w http.ResponseWriter, req *http.Request) {
	u := getUserData(w, req)
	if !alreadyLoggedIn(w, req) {
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}
	if u.Role != "007" {
		http.Error(w, "You must be 007 to enter the bar", http.StatusForbidden)
		return
	}
	showSessions() // for demonstration purposes
	tpl.ExecuteTemplate(w, "bar.gohtml", u)
}
