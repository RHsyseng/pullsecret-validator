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
	/*bs := []byte(`{"auths": {
		"assisted-ipv6-disconnecter000": {
			"auth": "ZHVtbXk6ZHVtbXk",
			"email": "jhendrix@karmalabs.com"
		},
		"bm-cluster-1-hyper.e2e.bos.redhat.com5000": {
			"auth": "a25pOmtuaQ=="
		}
	}}`)*/
	mux := http.NewServeMux()
	mux.HandleFunc("/", web.HandleRequest)
	fileServer := http.FileServer(http.Dir("./webmodule/"))
	mux.Handle("/webmodule/", http.StripPrefix("/webmodule", fileServer))
	err := http.ListenAndServe(":80", mux) // setting listening port
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
