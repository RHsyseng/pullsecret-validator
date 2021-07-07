package webmodule

import (
	"fmt"
	v "github.com/RHsyseng/lib-ps-validator"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func HandleRequest(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path != "/" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}

	if r.Method != "POST" {
		body := v.WebData{"Input your pull secret in json format", nil, nil, nil}
		t, err := template.ParseFiles("webmodule/web.html")
		fmt.Println(os.Getwd())

		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		t.Execute(w, body)

	} else {
		r.ParseForm()
		resData := v.Validate([]byte(r.FormValue("pullsecret")))

		body := v.WebData{r.FormValue("pullsecret"), resData.ResultOK, resData.ResultKO, resData.ResultCon}
		t, err := template.ParseFiles("webmodule/web.html")

		fmt.Println(os.Getwd())

		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		t.Execute(w, body)

	}
}
