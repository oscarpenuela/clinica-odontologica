package main

import (
	"database/sql"
	"log"

	_ "github.com/jmrobles/h2go"
)

func main() {
	conn, err := sql.Open("h2", "h2://sa@localhost/db?mem=true")
	if err != nil {
		log.Fatalf("Can't connet to H2 Database: %s", err)
	}
    err = conn.Ping()
    if err != nil {
        log.Fatalf("Can't ping to H2 Database: %s", err)
    }
    log.Printf("H2 Database connected")

    conn.Close()
}