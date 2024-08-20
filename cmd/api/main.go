package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	user     = "postgres"
	password = "1234"
	dbname   = "postgres"
	host     = "localhost"
	port     = 5432
)

func main() {
	// Connection to database
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	// Open the database
	db, err := sql.Open("postgres", psqlconn)
	if err != nil {
		panic(err)
	}
	fmt.Println("Database Connected!")
	// Close the database
	defer db.Close()
	insert := `insert into "icecream"("Name", "Price") values($1, $2)`
	// insert new raw to database table
	_, err = db.Exec(insert, "Vanilla", 4)
	if err != nil {
		panic(err)
	}
	_, err = db.Exec(insert, "Chocolate", 5)
	if err != nil {
		panic(err)
	}
	fmt.Println("Writing to database completed!")
}
