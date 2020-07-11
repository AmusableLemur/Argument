package main

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

// Post is a new thread on the discussion board
type Post struct {
	ID          int
	Title       string
	DateCreated string
}

// Connect is the function to prepare the database connection
func Connect(URI string) {
	var err error
	db, err = sql.Open("mysql", config.Database.URI)

	if err != nil {
		panic(err)
	}
}

// GetPosts loads all the available posts from the database
func GetPosts() []Post {
	results, err := db.Query("SELECT id, title, created_at FROM posts")

	if err != nil {
		log.Fatal("Something is wrong with the database structure")
	}

	var posts []Post

	for results.Next() {
		var post Post

		err = results.Scan(&post.ID, &post.Title, &post.DateCreated)

		if err != nil {
			log.Fatal(err.Error())
		}

		posts = append(posts, post)
	}

	return posts
}
