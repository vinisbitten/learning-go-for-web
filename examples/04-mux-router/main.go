package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"github.com/gorilla/mux"
	_ "github.com/go-sql-driver/mysql"
)

type Post struct {
	Id    int
	Title string
	Body  string
}

var db, err = sql.Open("mysql", "root:rootroot@tcp(127.0.0.1:3306)/go_course")

func main() {
	err := db.Ping()
	CheckErr(err)

	defer db.Close()

	r := mux.NewRouter()

	r.HandleFunc("/", HomeHandler)

	r.HandleFunc("/posts", ListPosts)

	r.HandleFunc("/post/{id}", ViewPost)

	r.HandleFunc("/create", CreatePost)

	r.HandleFunc("/update", UpdatePost)

	r.HandleFunc("/delete", DeletePost)

	log.Println("Listening on --> localhost:8080 <--")

	log.Fatal(http.ListenAndServe(":8080", r))
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	page := Post{Id: 1, Title: "Hello World", Body: "This is a sample page"}

	t := template.Must(template.ParseFiles("template/index.html"))
	if err := t.ExecuteTemplate(w, "index.html", page); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func ListPosts(w http.ResponseWriter, r *http.Request) {
	var posts []Post

	rows, err := db.Query("SELECT * FROM posts")
	CheckErr(err)

	for rows.Next() {
		var post Post
		rows.Scan(&post.Id, &post.Title, &post.Body)
		posts = append(posts, post)
	}

	t := template.Must(template.ParseFiles("template/posts.html"))
	if err := t.ExecuteTemplate(w, "posts.html", posts); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func ViewPost(w http.ResponseWriter, r *http.Request) {
	var post Post

	id := mux.Vars(r)["id"]

	row := db.QueryRow("SELECT * FROM posts WHERE id = ?", id)
	
	err := row.Scan(&post.Id, &post.Title, &post.Body)
	CheckErr(err)

	t := template.Must(template.ParseFiles("template/post.html"))
	if err := t.ExecuteTemplate(w, "post.html", post); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func CreatePost(w http.ResponseWriter, r *http.Request) {
	var post Post

	err := json.NewDecoder(r.Body).Decode(&post)
	CheckErr(err)

	stmt, err := db.Prepare("INSERT INTO posts(title, body) VALUES(?, ?)")
	CheckErr(err)

	result, err := stmt.Exec(post.Title, post.Body)
	CheckErr(err)

	rows, err := result.RowsAffected()
	CheckErr(err)

	if rows == 0 {
		fmt.Fprintf(w, "No rows created")
	} else {
		fmt.Fprintf(w, "Post created")
	}
}

func UpdatePost(w http.ResponseWriter, r *http.Request) {
	var post Post

	err := json.NewDecoder(r.Body).Decode(&post)
	CheckErr(err)

	stmt, err := db.Prepare("UPDATE posts SET title = ?, body = ? WHERE id = ?")
	CheckErr(err)

	result, err := stmt.Exec(post.Title, post.Body, post.Id)
	CheckErr(err)

	rows, err := result.RowsAffected()
	CheckErr(err)

	if rows == 0 {
		fmt.Fprintf(w, "No rows updated")
	} else {
		fmt.Fprintf(w, "Post updated")
	}
}

func DeletePost(w http.ResponseWriter, r *http.Request) {
	var post Post

	err := json.NewDecoder(r.Body).Decode(&post)
	CheckErr(err)

	stmt, err := db.Prepare("DELETE FROM posts WHERE id = ?")
	CheckErr(err)

	result, err := stmt.Exec(post.Id)
	CheckErr(err)

	rows, err := result.RowsAffected()
	CheckErr(err)

	if rows == 0 {
		fmt.Fprintf(w, "No rows deleted")
	} else {
		fmt.Fprintf(w, "Post deleted")
	}
}

func CheckErr(error) {
	if err != nil {
		panic(err)
	}
}
