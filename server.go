package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "abc123"
	dbname   = "testdb"
)

func Fibonacci(n uint64, db *sql.DB) uint64 {
	var myValue uint64
	userSql := "SELECT value from newfib WHERE mykey = $1"
	//check to see if this entry is in the table
	err := db.QueryRow(userSql, n).Scan(&myValue)
	if err == nil {
		//in the table, return immediately
		return myValue
	} else {
		//not in the table, need to create new entry
		fmt.Printf("Failed to get entry directly:%v\n", err)
		value := Fibonacci(n-1, db) + Fibonacci(n-2, db)
		insertDynStmt := `insert into "newfib"("mykey", "value") values($1, $2)`
		_, e := db.Exec(insertDynStmt, n, value)
		CheckError(e)
		return value
	}
}

//This is a simple debugging panic for error
func CheckError(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	// connection string
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	// open database
	db, err := sql.Open("postgres", psqlconn)
	CheckError(err)

	// close database
	defer db.Close()

	// check db
	err = db.Ping()
	CheckError(err)
	fmt.Println("Connected!")

	deleteStmt := `delete from newfib`
	_, e := db.Exec(deleteStmt)
	CheckError(e)

	//we store the first two elements of Fib series as a start
	insertDynStmt := `insert into "newfib"("mykey", "value") values($1, $2)`
	_, e = db.Exec(insertDynStmt, 0, 0)
	CheckError(e)
	_, e = db.Exec(insertDynStmt, 1, 1)
	CheckError(e)

	//os.Exit(1)

	r := mux.NewRouter()

	//API to generate Fib number
	r.HandleFunc("/fibonacci/{number}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		number := vars["number"]
		u, _ := strconv.ParseUint(number, 0, 64)
		fib := Fibonacci(u, db)
		fmt.Fprintf(w, "%v", fib)
	})

	//API to report the Fib entry number within a range
	r.HandleFunc("/order/{number}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		number := vars["number"]
		u, _ := strconv.ParseUint(number, 0, 64)
		var myValue uint64
		userSql := "SELECT count(*) from newfib WHERE value < $1"
		err := db.QueryRow(userSql, u).Scan(&myValue)
		if err == nil {
			fmt.Fprintf(w, "%v", myValue)
		}
	})

	//API to clear all Fib entry in table
	r.HandleFunc("/clear", func(w http.ResponseWriter, r *http.Request) {
		deleteStmt := `delete from newfib`
		_, e := db.Exec(deleteStmt)
		CheckError(e)
		fmt.Fprintf(w, "ok")
	})

	http.ListenAndServe(":8001", r)
}
