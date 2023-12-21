package main

import (
	"net/http"
	"text/template"
)

var tmp *template.Template

type Student struct {
	Name string
	Age  bool
}

var data []Student

func main() {
	http.HandleFunc("/", hello)
	tmp, _ = template.ParseGlob("templetes/*.html")
	data = []Student{{"Subham", true}, {"Jonsey", false}, {"Desodemona", true}, {"Peely", false}}
	http.ListenAndServe("", nil)
}

func hello(w http.ResponseWriter, r *http.Request) {
	tmp.ExecuteTemplate(w, "index.html", data)
}
