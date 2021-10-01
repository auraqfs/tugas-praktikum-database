package config

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func CreateConnection() *sql.DB {
	//load .env file

	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	// buka koneksi ke db

	db, err := sql.Open("postgres", os.Getenv("POSTGRES_URL"))

	if err != nil {
		panic(err)
	}

	// cek koneksi
	err = db.Ping()

	if err != nil {
		panic(err)
	}

	fmt.Println("Sukes connect ke db!")
	// return connection
	return db
}
