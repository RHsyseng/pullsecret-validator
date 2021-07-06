package main

import (
	"html/template"
	"log"
	"net/http"
)

type Data struct {
	Input  interface{}
	Result interface{}
}

func Validate(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path != "/" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}

	if r.Method == "GET" {
		body := Data{"Input your pull secret in json format", nil}
		t, _ := template.ParseFiles("web.html")
		t.Execute(w, body)

	} else {
		r.ParseForm()
		body := Data{r.FormValue("pullsecret"), r.FormValue("pullsecret")}

		t, _ := template.ParseFiles("web.html")
		t.Execute(w, body)

	}
}

func main() {

	http.HandleFunc("/", Validate)
	err := http.ListenAndServe(":80", nil) // setting listening port
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
