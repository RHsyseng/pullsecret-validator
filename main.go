package main

import (
	"html/template"
	"log"
	"net/http"
	web "github.com/RHsyseng/pullsecret-validator/webmodule"
)

type Data struct {
	Input  interface{}
	Result interface{}
}



func validatePS (input string) string{

}


func main() {

	http.HandleFunc("/", web.HandleRequest)
	err := http.ListenAndServe(":80", nil) // setting listening port
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
