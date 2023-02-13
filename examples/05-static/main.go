package main

import (
	"html/template"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Post struct {
	Id    int
	Title string
	Body  string
}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/", HomeHandler)

	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("static/"))))

	log.Println("Listening on --> localhost:8080 <--")
	log.Println("Css file at  --> localhost:8080/static/css/style.css <--")
	log.Fatal(http.ListenAndServe(":8080", r))
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	page := Post{Id: 1, Title: "Go page", Body: "This is a sample page"}

	t := template.Must(template.ParseFiles("template/index.html"))
	if err := t.ExecuteTemplate(w, "index.html", page); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
