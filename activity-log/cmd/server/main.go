package main

import (
	"log"

	"github.com/earthly/cloud-services-example/activity-log/internal/server"
)

func main() {
	log.Println("Starting listening on port 8080")
	srv := server.NewHTTPServer(":8080")
	log.Fatal(srv.ListenAndServe())
}
