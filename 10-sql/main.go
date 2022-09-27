package main

import (
	"database/sql"
	"fmt"
	"io"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// link := "root:P@ssw0rd@tcp(localhost:3306)/demo?charset=utf8"
	link := "golang:password@tcp(localhost:3306)/demo?charset=utf8"
	db, err := sql.Open("mysql", link)
	check(err)
	defer db.Close()
	err = db.Ping()
	check(err)

	http.Handle("/", http.HandlerFunc(index))
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8089", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "successfully connect to mysql")
}

func check(err error) {
	if err != nil {
		fmt.Println(err)
	}
}
