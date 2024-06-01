package main

import (
	"html/template"
	"log"
	"net/http"
)

const workingDir = "/Users/mitya/Desktop/mitya"

func index(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles(workingDir + "/views/index.html")

	if r.URL.Path != "/" {
		errorHandler(w, r)
		return
	}

	if err != nil {
		log.Fatal(err)
	}

	tmpl.Execute(w, nil)
}

func about(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles(workingDir + "/views/about.html")

	if r.URL.Path != "/about" {
		errorHandler(w, r)
		return
	}

	if err != nil {
		log.Fatal(err)
	}

	tmpl.Execute(w, nil)
}

func errorHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(404)
	if 404 == http.StatusNotFound {
		tmpl, err := template.ParseFiles(workingDir + "/views/404.html")
		if err != nil {
			log.Fatal(err)
		}

		tmpl.Execute(w, nil)
	}
}

func run() error {
	http.HandleFunc("/", index)
	http.HandleFunc("/about", about)
	return http.ListenAndServe(":8080", nil)
}

func main() {
	log.Fatal(run())
}
