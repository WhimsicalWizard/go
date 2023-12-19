package main

import (
	"net/http"
	"text/template"
)

var tmp *template.Template

func main() {
	tmp, _ = template.ParseFiles("index.html")
	http.HandleFunc("/", helloWord)
	http.ListenAndServe(":8080", nil)
}

func helloWord(w http.ResponseWriter, r *http.Request) {
	name := "Slone"
	tmp.Execute(w, name)
	//fmt.Fprint(w, "hello")
}
