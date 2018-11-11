package www

import (
	"html/template"
	"log"
	"net/http"
)

func Index(w http.ResponseWriter, _ *http.Request) {
	tpl := template.New("index.html")
	var err error
	tpl, err = tpl.ParseFiles("template/index.html")
	if err != nil {
		log.Printf("parse template error. %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(http.StatusText(http.StatusInternalServerError)))
		return
	}
  
  type Fields struct {
		Title string
		Short_url string
	}
	var fields = Fields{"yb.gd", "https://yb.gd/U"}

	err = tpl.Execute(w, fields)
	if err != nil {
		log.Printf("execute template error. %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(http.StatusText(http.StatusInternalServerError)))
	}
}
