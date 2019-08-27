package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	a := App{}
	a.Initialize(
		os.Getenv("APP_DB_USERNAME"),
		os.Getenv("APP_DB_PASSWORD"),
		os.Getenv("APP_DB_NAME"),
		os.Getenv("SSLMODE"))

	//a.Run(":8080")
	log.Printf("Start Server")
	log.Fatal(http.ListenAndServe(":8000", a.Router))

}
