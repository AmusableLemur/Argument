package app

import (
	"html/template"
	"net/http"

	"github.com/AmusableLemur/Argument/internal/config"
	"github.com/AmusableLemur/Argument/internal/database"
	"github.com/gorilla/mux"
)

// PostIndex is used for the individual post pages
type PostIndex struct {
	PageTitle string
	Post      database.Post
}

// PostsIndex is used for the index page
type PostsIndex struct {
	PageTitle string
	Posts     []database.Post
}

// SetupHandler prepares the route handler
func SetupHandler() *mux.Router {
	r := mux.NewRouter()
	t := template.Must(template.ParseGlob("templates/*.tmpl"))

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		err := t.ExecuteTemplate(w, "index.tmpl", PostsIndex{
			PageTitle: config.Conf.Title,
			Posts:     database.GetPosts(),
		})

		if err != nil {
			panic(err)
		}
	})

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		p := database.Post{Title: r.FormValue("title")}

		database.SavePost(p)

		err := t.ExecuteTemplate(w, "index.tmpl", PostsIndex{
			PageTitle: config.Conf.Title,
			Posts:     database.GetPosts(),
		})

		if err != nil {
			panic(err)
		}
	}).Methods("POST")

	return r
}
