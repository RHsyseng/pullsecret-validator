package main

import (
	web "github.com/RHsyseng/pullsecret-validator/webmodule"
	"log"
	"net/http"
)

type Data struct {
	Input  interface{}
	Result interface{}
}

func main() {

	http.HandleFunc("/", web.HandleRequest)
	err := http.ListenAndServe(":80", nil) // setting listening port
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
