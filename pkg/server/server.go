package server

import (
    "net/http"
    "fmt"
)

func StartServer() {
    if err := http.ListenAndServe(":8080", nil); err != nil {
        fmt.Println("Error starting server")
    }
}
