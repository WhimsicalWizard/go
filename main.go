package main

import (
	"net/http"
	"text/template"
)

var tmp *template.Template

func main() {
	http.HandleFunc("/", hello)
	http.HandleFunc("/processForm", processHandler)
	tmp, _ = template.ParseGlob("templetes/*.html")

	http.ListenAndServe("", nil)
}

func hello(w http.ResponseWriter, r *http.Request) {
	tmp.ExecuteTemplate(w, "index.html", nil)
}

func processHandler(w http.ResponseWriter, r *http.Request) {
	n := r.FormValue("name")
	tmp.ExecuteTemplate(w, "form.html", n)
}
