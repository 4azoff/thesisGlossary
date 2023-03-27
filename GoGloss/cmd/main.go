package main

import (
	"html/template"
	"net/http"
)

type Term struct {
	Name        string
	Description string
}

var terms = []Term{
	{Name: "Golang", Description: "A statically typed, compiled programming language designed at Google"},
	{Name: "HTML", Description: "A markup language used for creating web pages"},
	{Name: "CSS", Description: "A style sheet language used for describing the presentation of a document written in HTML"},
}

func renderTerms(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("terms.html"))
	err := tmpl.Execute(w, terms)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func renderMindMap(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("mindMap.html"))
	err := tmpl.Execute(w, terms)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func main() {

	http.HandleFunc("/termsGo", renderTerms)
	http.HandleFunc("/mindMap", renderMindMap)

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	http.ListenAndServe(":8080", nil)
}
