package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

// Version control: Use Git for version control and organize your project into meaningful commits.
// Documentation: Document your code with comments and README files.
// Error handling: Always handle errors properly in Go using the if err != nil pattern.
// Logging: Use Go’s standard log package for logging important information.
// Environment variables: Use environment variables for sensitive information (like database passwords) and configuration settings.

// creating an db connection
func db() *sql.DB {
	db, err := sql.Open("mysql", "user:password@tcp(127.0.0.1:3306)/dbname")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	return db
}

// creating an handler
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, Go HTTP Server!")
}

func main() {
	db := db()
	println(db)
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
