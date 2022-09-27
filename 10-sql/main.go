package main

import (
	"database/sql"
	"fmt"
	"io"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB
var err error

func main() {
	// link := "root:P@ssw0rd@tcp(localhost:3306)/demo?charset=utf8"
	link := "golang:password@tcp(localhost:3306)/demo?charset=utf8"
	db, err = sql.Open("mysql", link)
	check(err)
	defer db.Close()
	err = db.Ping()
	check(err)

	http.Handle("/", http.HandlerFunc(index))
	http.Handle("/create", http.HandlerFunc(create))
	http.Handle("/amigos", http.HandlerFunc(amigos))
	http.Handle("/favicon.ico", http.NotFoundHandler())
	fmt.Println("Server starts on port 8089")
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

func amigos(w http.ResponseWriter, r *http.Request) {
	rows, err := db.Query(`SELECT aName FROM demo`)
	check(err)
	defer rows.Close()

	var s, name string
	s = "RETRIEVED RECORDS:\n"

	for rows.Next() {
		err = rows.Scan(&name)
		check(err)
		s += name + "\n"
	}

	fmt.Fprintln(w, s)
}

func create(w http.ResponseWriter, r *http.Request) {
	stmt, err := db.Prepare(`CREATE TABLE customer (name VARCHAR(20))`)
	check(err)
	defer stmt.Close()

	re, err := stmt.Exec()
	check(err)

	n, err := re.RowsAffected()
	check(err)

	fmt.Fprintln(w, "CREATED TABLE customer", n)
}
