package server

import (
    "net/http"
    "path/filepath"
    "html/template"
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
