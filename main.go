package main

import (
	"github.com/yigitsadic/gocandela/client"
	"github.com/yigitsadic/gocandela/handlers"
	"log"
	"net/http"
)

func main() {
	c := client.NewHTTPClient()

	http.HandleFunc("/", handlers.EarthquakeHandler(c))

	log.Println("Server started on port 8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("unable to listen on port 8080")
	}
}
