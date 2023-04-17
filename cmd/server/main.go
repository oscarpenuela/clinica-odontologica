package main

import (
	"database/sql"
	"log"

	"github.com/gin-gonic/gin"
	_ "github.com/jmrobles/h2go"
	"github.com/oscarpenuela/clinica-odontologica.git/cmd/server/handler"
	"github.com/oscarpenuela/clinica-odontologica.git/internal/dentist"
	"github.com/oscarpenuela/clinica-odontologica.git/pkg/store"
)

func main() {
	db, err := sql.Open("h2", "h2://sa@localhost/db?mem=true")
	if err != nil {
		log.Fatalf("Can't connect to H2 Database: %s", err)
        panic(err)
	}
    err = db.Ping()
    if err != nil {
        log.Fatalf("Can't ping to H2 Database: %s", err)
    }
    log.Printf("H2 Database connected")
    
    _, err = db.Exec("CREATE DATABASE clinica")
	if err != nil {
		log.Fatal(err)
	}
    

    _, err = db.Exec(`
        CREATE TABLE IF NOT EXISTS dentists (
            id INTEGER PRIMARY KEY AUTOINCREMENT,
            apellido TEXT NOT NULL,
            nombre TEXT NOT NULL,
            matricula TEXT NOT NULL
        )
    `)
    if err != nil {
        panic(err)
    }

    _, err = db.Exec(`
        CREATE TABLE IF NOT EXISTS patients (
            id INTEGER PRIMARY KEY AUTOINCREMENT,
            nombre TEXT NOT NULL,
            apellido TEXT NOT NULL,
            domicilio TEXT NOT NULL,
            dni TEXT NOT NULL,
            fecha_de_alta TEXT NOT NULL
        )
    `)
    if err != nil {
        panic(err)
    }

    _, err = db.Exec(`
        CREATE TABLE IF NOT EXISTS appointments (
            id INTEGER PRIMARY KEY AUTOINCREMENT,
            dentist_id INTEGER NOT NULL,
            patient_id INTEGER NOT NULL,
            date TEXT NOT NULL,
            time TEXT NOT NULL,
            description TEXT NOT NULL,
            FOREIGN KEY (dentist_id) REFERENCES dentists(id),
            FOREIGN KEY (patient_id) REFERENCES patients(id)
        )
    `)
    if err != nil {
        panic(err)
    }

    storage := store.SqlStore{db}
    repo := dentist.Repository{&storage}
    service := dentist.Service{&repo}
    dentistHandler := handler.DentistHandler{&service}

    r:= gin.Default()
    dentists := r.Group("dentists")
    {dentists.GET(":id", dentistHandler.GetById)}

    r.Run(":8080")
}