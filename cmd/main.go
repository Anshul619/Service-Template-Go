package main

import (
	"log"
	"net/http"
	v1 "template-service/internal/handler/v1"
)

func main() {

	v1.InitRoutes()

	log.Println("Listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))

}
