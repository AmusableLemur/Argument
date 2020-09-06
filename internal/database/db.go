package database

import (
	"database/sql"

	"github.com/AmusableLemur/Argument/internal/config"

	// Imported to get MySQL support
	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

// Post is a new thread on the discussion board
type Post struct {
	ID          int
	Title       string
	DateCreated string
}

func init() {
	Connect(config.Conf.Database.URI)
}

// Connect prepares the database connection
func Connect(URI string) {
	if db != nil {
		return
	}

	var err error
	db, err = sql.Open("mysql", URI)

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
	if db != nil {
		db.Close()
		db = nil
	}
}

// GetPosts loads all the available posts from the database
func GetPosts() []Post {
	results, _ := db.Query("SELECT id, title, created_at FROM posts")

	var posts []Post

	for results.Next() {
		var post Post
		results.Scan(&post.ID, &post.Title, &post.DateCreated)
		posts = append(posts, post)
	}

	return posts
}

// SavePost persists a given post to database
func SavePost(post Post) (int64, error) {
	query, _ := db.Prepare("INSERT INTO posts(title, created_at, created_by) values (?, NOW(), 0)")

	result, err := query.Exec(post.Title)

	if err != nil {
		id, _ := result.LastInsertId()

		return id, nil
	}

	return 0, err
}
