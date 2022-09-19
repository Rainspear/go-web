package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
)

func alreadyLoggedIn(w http.ResponseWriter, r *http.Request) bool {
	c, err := r.Cookie("session")
	if err != nil {
		return false
	}

	sess, ok := dbSessions[c.Value]
	if !ok {
		return false
	}

	sess.lastActivity = time.Now() // set time again when resignin
	c.MaxAge = sessionLength
	http.SetCookie(w, c)
	return true
}

// return user data, if not create a new session and return new data user
func getUserData(w http.ResponseWriter, r *http.Request) user {
	c, err := r.Cookie("session")
	if err != nil {
		c = &http.Cookie{
			Name:  "session",
			Value: uuid.New().String(),
		}
	}
	c.MaxAge = sessionLength
	http.SetCookie(w, c)

	var u user
	if s, ok := dbSessions[c.Value]; ok {
		s.lastActivity = time.Now()
		dbSessions[c.Value] = s
		u = dbUsers[s.userId]
	}
	return u
}

func showSessions() {
	// fmt.Println("********")
	for k, v := range dbSessions {
		fmt.Println(k, v.userId)
	}
	fmt.Println("")
}

func cleanSessions() {
	fmt.Println("BEFORE CLEAN") // for demonstration purposes
	showSessions()              // for demonstration purposes
	for k, v := range dbSessions {
		if time.Now().Sub(v.lastActivity) > (time.Second * 30) {
			delete(dbSessions, k)
		}
	}
	dbSessionsCleaned = time.Now()
	fmt.Println("AFTER CLEAN") // for demonstration purposes
	showSessions()             // for demonstration purposes
}
