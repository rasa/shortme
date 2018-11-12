package www

import (
  "bytes"
	"html/template"
	"log"
	"net/http"

	"github.com/rasa/shortme/conf"
)

var bb bytes.Buffer

func Init() {
	tpl := template.New("index.html")
	var err error
	tpl, err = tpl.ParseFiles("template/index.html")
	if err != nil {
		log.Fatalf("parse template error. %v", err)
	}

	err = tpl.Execute(&bb, &conf.Conf.Common)
	if err != nil {
		log.Fatalf("execute template error. %v", err)
	}
}

func Index(w http.ResponseWriter, _ *http.Request) {
  if bb.Len() == 0 {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(http.StatusText(http.StatusInternalServerError)))
		return
  }
  w.Write(bb.Bytes())
}
