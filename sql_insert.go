package main

import (
	"database/sql"

	"fmt"

	"log"

	_ "github.com/go-sql-driver/mysql"
)

// NewDB function
func NewDB() *sql.DB {
	db, _ := sql.Open("mysql", "root:iamaprogrammer@tcp(127.0.0.1:3306)/notesapp")
	err := db.Ping()
	if err != nil {
		panic(err)
	}
	return db
}

func main() {
	db := NewDB()

	// res, err := db.Exec(
		"INSERT INTO users(username, email, password, bio) VALUES(?, ?, ?, ?)",
		"lorem",
		"lorem@gmail.com",
		"password",
		"Bioooo",
	)
	if err != nil {
		log.Fatal(err)
	}
	last, _ := res.LastInsertId()
	affected, _ := res.RowsAffected()
	fmt.Println(last, affected)

	// USING PREPARED STATEMENTS
	stmt, err := db.Prepare("INSERT INTO users(username, email, password, bio) VALUES(?, ?, ?, ?)")
	if err != nil {
		log.Fatal(err)
	}
	res, err := stmt.Exec("lorem", "lorem@gmail.com", "password", "Bioooo")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(res.LastInsertId())

}
