package server

import (
    "encoding/json"
	"html/template"
	"net/http"
	"path/filepath"
    "strings"

	"github.com/Bialek328/go-casino-247/pkg/session"
)

var tmpl *template.Template

func InitTemplate() {
    tmpl = template.Must(template.ParseFiles(filepath.Join("..", "..", "ui", "templates", "index.html")))
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
    data := struct {
        Title string
        Heading string
        Message string
    }{
        Title: "NV Casino 247",
        Heading: "Hello",
        Message: "No hej przystojniaku",
    }
    if err := tmpl.ExecuteTemplate(w, "index.html", data); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }
} 

func listSessionsHandler(w http.ResponseWriter, r *http.Request) {
    ids := make([]string, 0, len(session.Sessions))
    for id := range(session.Sessions) {
        ids = append(ids, id)
    }
    response := map[string][]string {
        "Sessions": ids,
    }
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(response)
}

func addNewSession(w http.ResponseWriter, r *http.Request) {
    newSession := session.NewSession()
    session.Sessions[newSession.ID] = newSession
    
    // response := map[string]interface{}{
    //     "Sessions": session.Sessions,
    // }

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(session.Sessions)
}

func HandleSessionsEndpoint(w http.ResponseWriter, r *http.Request) {
    switch r.Method {
    case http.MethodGet:
        listSessionsHandler(w, r)
    case http.MethodPost:
        addNewSession(w, r)
    default:
        http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
    }
}

func HandleSingleSession(w http.ResponseWriter, r *http.Request) {
    sessionID, valid := extractIdFromUrl(r.URL.Path)
    if !valid {
        http.Error(w, "Invalid session ID", http.StatusBadRequest)
        return
    }

    switch r.Method {
    case http.MethodGet:
        currentSession, exist := session.Sessions[sessionID]
        if !exist {
            http.Error(w, "No such sesion", http.StatusNotFound)
            return
        }
        w.Header().Set("Content-Type", "application/json")
        json.NewEncoder(w).Encode(currentSession)
    }
}

func extractIdFromUrl(url string) (string, bool) {
    split := strings.Split(url, "/")
    if len(split) >= 5 && split[3] == "sessions" {
        return split[4], true
    } 
    return "", false
}

func HandlePlayerActionHit(w http.ResponseWriter, r *http.Request) {

}

func HandlePlayerActionStand(w http.ResponseWriter, r *http.Request) {

}

func HandleAddPlayer(w http.ResponseWriter, r *http.Request) {

}

func HandleRemovePlayer(w http.ResponseWriter, r *http.Request) {

}

func HandlePlayers(w http.ResponseWriter, r *http.Request) {

}

func HandleStartGame(w http.ResponseWriter, r *http.Request) {

}
