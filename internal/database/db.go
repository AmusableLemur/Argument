package database

import (
	"database/sql"

	"github.com/AmusableLemur/Argument/internal/config"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

// Post is a new thread on the discussion board
type Post struct {
	ID          int
	Title       string
	DateCreated string
}

// Connect prepares the database connection
func Connect(URI string) {
	if db != nil {
		return
	}

	var err error
	db, err = sql.Open("mysql", config.Conf.Database.URI)

	if err == nil {
		// Check that we actually managed to get a connection
		err = db.Ping()
	}

	if err != nil {
		panic(err)
	}
}

// Disconnect the database
func Disconnect() {
	if db == nil {
		return
	}

	db.Close()
	db = nil
}

// GetPosts loads all the available posts from the database
func GetPosts() []Post {
	Connect(config.Conf.Database.URI)

	results, _ := db.Query("SELECT id, title, created_at FROM posts")

	var posts []Post

	for results.Next() {
		var post Post
		results.Scan(&post.ID, &post.Title, &post.DateCreated)
		posts = append(posts, post)
	}

	return posts
}
