package server

import (
    "net/http"
)

func SetupRoutes() *http.ServeMux {
    mux := http.NewServeMux()
    // mux.HandleFunc("/v1/", )
    // mux.HandleFunc("/v1/blackjack/", ) // general info about blackjack
    mux.HandleFunc("/v1/blackjack/sessions", HandleSessionsEndpoint) //list all active sessions
    mux.HandleFunc("/v1/blackjack/sessions/", HandleSingleSession) // make this web socket
    mux.HandleFunc("/v1/blackjack/sessions/{uuid}/add_player", HandleAddPlayer)
    mux.HandleFunc("v1/blackjack/sessions/{uuid}/remove_player", HandleRemovePlayer)
    mux.HandleFunc("v1/blackjack/sessions/{uuid}/start_game", HandleStartGame)
    mux.HandleFunc("v1/blackjack/sessions/{uuid}/players", HandlePlayers)
    mux.HandleFunc("v1/blackjack/sessions/{uuid}/players/uuid/hit", HandlePlayerActionHit)
    mux.HandleFunc("v1/blackjack/sessions/{uuid}/players/uuid/stand", HandlePlayerActionStand)
    return mux
}
