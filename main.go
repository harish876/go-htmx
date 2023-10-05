package main

import (
	"fmt"
	db "go-htmx/db"
	"log"
	"net/http"
	"strings"
	"text/template"
)

type Film struct {
	Title    string
	Director string
}

func main() {
	database, err := db.Init()
	if err != nil {
		log.Fatal(err)
	}

	// h1 := func(w http.ResponseWriter, r *http.Request) {
	// 	films := make(map[string][]Film)
	// 	tmpl := template.Must(template.ParseFiles("templates/index.html"))
	// 	rows, _ := database.Query("SELECT title,Director from films")

	// 	for rows.Next() {
	// 		var film Film
	// 		rows.Scan(&film.Title, &film.Director)
	// 		films["films"] = append(films["films"], film)
	// 	}

	// 	tmpl.Execute(w, films)
	// }

	h2 := func(w http.ResponseWriter, r *http.Request) {
		Title := r.PostFormValue("title")
		Director := r.PostFormValue("director")

		newFilm := Film{
			Title:    Title,
			Director: Director,
		}

		statement, _ := database.Prepare("INSERT INTO films (title,director) VALUES (?,?)")
		statement.Exec(newFilm.Title, newFilm.Director)

		films := make(map[string][]Film)
		rows, _ := database.Query("SELECT title,Director from films")
		tmpl := template.Must(template.ParseFiles("templates/examples.html"))
		for rows.Next() {
			var film Film
			rows.Scan(&film.Title, &film.Director)
			films["films"] = append(films["films"], film)
		}

		tmpl.Execute(w, films)
	}

	h3 := func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path
		prefix := "/delete-film/"
		if strings.HasPrefix(path, prefix) {
			title := strings.TrimPrefix(path, prefix)
			statement, _ := database.Prepare("DELETE FROM films WHERE title=$1")
			statement.Exec(title)
			log.Printf("Deleting item with title: %s", title)
		} else {
			http.NotFound(w, r)
			fmt.Println("Not found")
		}
	}

	h4 := func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles("templates/modal.html"))
		segments := strings.Split(r.URL.Path, "/")
		log.Println(segments)

		if len(segments) < 3 {
			http.NotFound(w, r)
			return
		}

		title := segments[2]
		director := segments[3]

		editFilm := Film{
			Title:    title,
			Director: director,
		}

		tmpl.Execute(w, editFilm)
	}

	h5 := func(w http.ResponseWriter, r *http.Request) {
		films := make(map[string][]Film)
		rows, _ := database.Query("SELECT title,Director from films")
		tmpl := template.Must(template.ParseFiles("templates/examples.html"))
		for rows.Next() {
			var film Film
			rows.Scan(&film.Title, &film.Director)
			films["films"] = append(films["films"], film)
		}

		tmpl.Execute(w, films)
	}

	h6 := func(w http.ResponseWriter, r *http.Request) {
		Title := r.PostFormValue("title")
		Director := r.PostFormValue("director")

		if Title != "" || Director != "" {
			statement, _ := database.Prepare("UPDATE films SET director=? WHERE title=?")
			statement.Exec(Director, Title)
		}

		films := make(map[string][]Film)
		rows, _ := database.Query("SELECT title,Director from films")
		tmpl := template.Must(template.ParseFiles("templates/examples.html"))
		for rows.Next() {
			var film Film
			rows.Scan(&film.Title, &film.Director)
			films["films"] = append(films["films"], film)
		}

		tmpl.Execute(w, films)
	}

	h7 := func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles("templates/add-film-modal.html"))
		tmpl.Execute(w, nil)
	}

	http.HandleFunc("/", h5)
	http.HandleFunc("/add-film/", h2)
	http.HandleFunc("/add-film-modal/", h7)
	http.HandleFunc("/delete-film/", h3)
	http.HandleFunc("/uikit-modal/", h4)
	// http.HandleFunc("/examples/", h5)
	http.HandleFunc("/edit-film/", h6)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
