package main

import (
	web "github.com/RHsyseng/pullsecret-validator/webmodule"
	"log"
	"net/http"
)

func main() {

	mux := http.NewServeMux()
	mux.HandleFunc("/", web.HandleRequest)
	fileServer := http.FileServer(http.Dir("./webmodule/"))
	mux.Handle("/webmodule/", http.StripPrefix("/webmodule", fileServer))
	err := http.ListenAndServe(":8080", mux) // setting listening port
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}

}
