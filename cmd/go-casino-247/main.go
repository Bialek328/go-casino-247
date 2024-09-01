package main

import "github.com/Bialek328/go-casino-247/pkg/server"


func main() {
    server.InitTemplate()
    server.SetupRoutes()

    server.StartServer()
}
