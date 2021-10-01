package main

import (
	"go-postgres/config"
	"go-postgres/router"
	"log"
	"net/http"
)

func main() {
	config.CreateConnection()

	r := router.Router()

	log.Fatal(http.ListenAndServe(":9000", r))

}
