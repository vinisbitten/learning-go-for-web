package main

import (
	"html/template"
	"log"
	"net/http"
)

type Post struct {
	Id    int
	Title string
	Body  string
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		// Using curly braces in html template you can access the fields of the struct
		// {{.Id}} {{.Title}} {{.Body}} --> chech the index.html file
		post := Post{Id: 1, Title: "Hello World", Body: "This is a sample post"}

		t := template.Must(template.ParseFiles("template/index.html"))
		if err := t.ExecuteTemplate(w, "index.html", post); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})
	log.Println("Listening on --> localhost:8080 <--")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
