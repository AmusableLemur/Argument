package app

import (
	"html/template"
	"net/http"

	"github.com/AmusableLemur/Argument/internal/auth"
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
	t := template.Must(template.ParseGlob(config.Conf.Root + "templates/*.tmpl"))

	// Index page
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		t.ExecuteTemplate(w, "index.tmpl", PostsIndex{
			PageTitle: config.Conf.Title,
			Posts:     database.GetPosts(),
		})
	})

	// Creating a post
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		p := database.Post{Title: r.FormValue("title")}

		database.SavePost(p)

		t.ExecuteTemplate(w, "index.tmpl", PostsIndex{
			PageTitle: config.Conf.Title,
			Posts:     database.GetPosts(),
		})
	}).Methods(http.MethodPost)

	// Registration
	r.HandleFunc("/register", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			p := auth.HashPassword(r.FormValue("password"))
			u := auth.NormalizeUsername(r.FormValue("username"))
			user := database.User{Username: u, Password: p}

			id, _ := database.CreateUser(user)

			if id > 0 {
				// Set the user as logged in
				// https://github.com/gorilla/sessions

				http.Redirect(w, r, "/", http.StatusTemporaryRedirect)

				return
			}

			// Show message about error logging in
		}

		t.ExecuteTemplate(w, "register.tmpl", PostsIndex{
			PageTitle: config.Conf.Title,
			Posts:     database.GetPosts(),
		})
	}).Methods(http.MethodGet, http.MethodPost)

	// Login
	r.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			p := auth.HashPassword(r.FormValue("password"))
			u := auth.NormalizeUsername(r.FormValue("username"))
			user := database.User{Username: u, Password: p}

			id, _ := database.FindUser(user)

			if id > 0 {
				// Set the user as logged in
				// https://github.com/gorilla/sessions

				http.Redirect(w, r, "/", http.StatusTemporaryRedirect)

				return
			}

			// Show message about user not found
		}

		t.ExecuteTemplate(w, "register.tmpl", PostsIndex{
			PageTitle: config.Conf.Title,
			Posts:     database.GetPosts(),
		})
	}).Methods(http.MethodGet, http.MethodPost)

	return r
}
