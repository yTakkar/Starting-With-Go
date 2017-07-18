package main

import (
	"database/sql"

	"log"

	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

// NewDB function
func NewDB() *sql.DB {
	db, _ := sql.Open("mysql", "root:iamaprogrammer@tcp(127.0.0.1:3306)/instagram")
	err := db.Ping()
	if err != nil {
		panic(err)
	}
	return db
}

func main() {
	db := NewDB()

	// FOR MULTIPLE ROWS
	rows, err := db.Query("SELECT id, username FROM users WHERE id=? AND username=?", 224, "coldplay")
	if err != nil {
		log.Fatal(err)
	}
	for rows.Next() {
		var (
			id       int
			username string
		)
		rows.Scan(&id, &username)
		fmt.Println(id, username)
	}
	er := rows.Err()
	if er != nil {
		log.Fatal(er)
	}

	// SELECT WITH PREPARED STATEMENTS
	stmt, err := db.Prepare("SELECT id, username FROM users WHERE id=? AND username=?")
	if err != nil {
		log.Fatal(err)
	}
	rows, err := stmt.Query(224, "coldplay")
	for rows.Next() {
		var (
			id       int
			username string
		)
		rows.Scan(&id, &username)
		fmt.Println(id, username)
	}

	// FOR SINGLE ROW
	var (
		id       int
		username string
	)
	row := db.QueryRow("SELECT id, username FROM users WHERE id=? AND username=?", 224, "coldplay").Scan(&id, &username)
	fmt.Println(id, username)

}
