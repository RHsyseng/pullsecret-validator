package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)



func validate(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method) //get request method
	if r.URL.Path != "/" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}

	if r.Method == "GET" {
		t, _ := template.ParseFiles("web.html")
		t.Execute(w, nil)
	} else {
		r.ParseForm()
		// logic part of log in
		fmt.Println("pull secret:", r.Form["pullsecret"])
		fmt.Fprintf(w, "Post from website! r.PostFrom = %v\n", r.PostForm)
		ps := r.FormValue("pullsecret")
		fmt.Fprintf(w, "Pull = %s\n", ps)


	}
}

func main() {

	http.HandleFunc("/", validate)
	err := http.ListenAndServe(":80", nil) // setting listening port
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
