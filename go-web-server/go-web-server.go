// go practice building a web server

package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"time"
)

type Film struct {
	Title    string
	Director string
}

func main() {
	fmt.Println("Building a web server!")

	//first handler function to display films described below
	h1 := func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles("index.html"))
		films := map[string][]Film{
			"Films": {
				{Title: "The Godfather", Director: "Francis Ford Coppola"},
				{Title: "Blade Runner", Director: "Ridley Scott"},
				{Title: "The Thing", Director: "John Carpenter"},
			},
		}
		tmpl.Execute(w, films)
	}

	//second handler function for inputing data into the form and displaying entries
	h2 := func(w http.ResponseWriter, r *http.Request) {
		time.Sleep((1 * time.Second))
		title := r.PostFormValue("title")
		director := r.PostFormValue("director")
		tmpl := template.Must(template.ParseFiles("index.html"))
		tmpl.ExecuteTemplate(w, "film-list-element", Film{Title: title, Director: director})
	}

	//handler functions for h1 and h2 functions built above
	http.HandleFunc("/", h1)
	http.HandleFunc("/add-film/", h2)

	//handle requests for port 8000
	log.Fatal(http.ListenAndServe(":8000", nil))
}
