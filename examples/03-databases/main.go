package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"

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

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		// Using curly braces in html template you can access the fields of the struct
		// {{.Id}} {{.Title}} {{.Body}} --> chech the index.html file
		post := Post{Id: 1, Title: "Hello World", Body: "This is a sample post"}

		t := template.Must(template.ParseFiles("template/index.html"))
		if err := t.ExecuteTemplate(w, "index.html", post); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	http.HandleFunc("/posts", func(w http.ResponseWriter, r *http.Request) {

		posts := Read()

		t := template.Must(template.ParseFiles("template/posts.html"))
		if err := t.ExecuteTemplate(w, "posts.html", posts); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	http.HandleFunc("/create", func(w http.ResponseWriter, r *http.Request) {
		var post Post

		err := json.NewDecoder(r.Body).Decode(&post)
		CheckErr(err)

		rows := Create(post.Title, post.Body)

		if rows == 0 {
			fmt.Fprintf(w, "No rows created")
		} else {
			fmt.Fprintf(w, "Post created")
		}
	})

	http.HandleFunc("/update", func(w http.ResponseWriter, r *http.Request) {
		var post Post

		err := json.NewDecoder(r.Body).Decode(&post)
		CheckErr(err)

		rows := Update(post.Id, post.Title, post.Body)

		if rows == 0 {
			fmt.Fprintf(w, "No rows updated")
		} else {
			fmt.Fprintf(w, "Post updated")
		}
	})

	http.HandleFunc("/delete", func(w http.ResponseWriter, r *http.Request) {
		var post Post

		err := json.NewDecoder(r.Body).Decode(&post)
		CheckErr(err)

		rows := Delete(post.Id)

		if rows == 0 {
			fmt.Fprintf(w, "No rows deleted")
		} else {
			fmt.Fprintf(w, "Post deleted")
		}
	})

	log.Println("Listening on --> localhost:8080 <--")

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func CheckErr(error) {
	if err != nil {
		panic(err)
	}
}

func Create(title string, body string) int64 {
	stmt, err := db.Prepare("INSERT INTO posts(title, body) VALUES(?, ?)")
	CheckErr(err)

	result, err := stmt.Exec(title, body)
	CheckErr(err)

	rows, err := result.RowsAffected()
	CheckErr(err)

	return rows
}

func Read() []Post {
	var posts []Post

	rows, err := db.Query("SELECT * FROM posts")
	CheckErr(err)

	for rows.Next() {
		var post Post
		rows.Scan(&post.Id, &post.Title, &post.Body)
		posts = append(posts, post)
	}

	return posts
}

func Update(id int, title string, body string) int64 {
	stmt, err := db.Prepare("UPDATE posts SET title = ?, body = ? WHERE id = ?")
	CheckErr(err)

	result, err := stmt.Exec(title, body, id)
	CheckErr(err)

	rows, err := result.RowsAffected()
	CheckErr(err)

	return rows
}

func Delete(id int) int64 {
	stmt, err := db.Prepare("DELETE FROM posts WHERE id = ?")
	CheckErr(err)

	result, err := stmt.Exec(id)
	CheckErr(err)

	rows, err := result.RowsAffected()
	CheckErr(err)

	return rows
}
