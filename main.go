package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
	"golang.org/x/net/context"
)

const (
	host   = "localhost"
	port   = 5432
	user   = "ccummer"
	dbname = "query_context"
)

func main() {
	fmt.Println("Running QueryContext test...")

	// Connect to the postgres db
	psqlInfo := fmt.Sprintf(
		"host=%s port=%d user=%s dbname=%s sslmode=disable",
		host,
		port,
		user,
		dbname,
	)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Executing query...")

	// Execute the query
	ids := []interface{}{2, 3, 4}
	rows, err := db.QueryContext(context.Background(), "SELECT id, data FROM experiment WHERE id IN ($1, $2, $3)", ids...)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	fmt.Println("Scanning rows...")

	// Spit out the result
	for rows.Next() {
		var id int
		var data string

		if err := rows.Scan(&id, &data); err != nil {
			log.Fatal(err)
		}

		fmt.Printf("%d\t%s\n", id, data)
	}

	if rows.Err() != nil {
		log.Fatal(rows.Err())
	}

	// Die
	fmt.Println("Done")
	os.Exit(0)
}
