package webmodule

import "net/http"

type Data struct {
	Input  interface{}
	Result interface{}
}

func HandleRequest(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path != "/" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}

	if r.Method != "POST" {
		body := Data{"Input your pull secret in json format", nil}
		t, _ := template.ParseFiles("web.html")
		t.Execute(w, body)

	} else {
		r.ParseForm()
		res := validatePS(r.FormValue("pullsecret"))

		body := Data{r.FormValue("pullsecret"), res}
		t, _ := template.ParseFiles("web.html")
		t.Execute(w, body)

	}
}

