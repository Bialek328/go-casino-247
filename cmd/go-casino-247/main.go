package main

import (
	"log"
	"net/http"

	"github.com/Bialek328/go-casino-247/pkg/server"
)


func main() {
    router := server.SetupRoutes()

    log.Fatal(http.ListenAndServe(":8080", router))
}
