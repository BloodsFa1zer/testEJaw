package main

import (
	"log"
	"net/http"
	"test/routes"
)

func main() {
	username := "admin"
	password := "password"

	mux := http.NewServeMux()
	routes.UserRoutes(mux, username, password)

	log.Println("Server is running on port 8080")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatal(err)
	}
}
