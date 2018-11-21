package www

import (
	"bytes"
	"html/template"
  "io/ioutil"
	"log"
	"net/http"

	"github.com/rasa/shortme/conf"
	_template "github.com/rasa/shortme/template"
)

const (
	html          = "/index.html"
	template_html = "template" + html
)

var bb bytes.Buffer

func Init() {
  fh, err := _template.Assets.Open(html)
  defer fh.Close()
  if err != nil {
    log.Fatalf("Failed to open %v: %v", template_html, err)
  }
  var data []byte
  data, err = ioutil.ReadAll(fh)
	if err != nil {
		log.Fatalf("Failed to read %v: %v", template_html, err)
	}
	tpl := template.New(html)
	tpl, err = tpl.Parse(string(data))
	if err != nil {
		log.Fatalf("Failed to parse %v: %v", template_html, err)
	}

	err = tpl.Execute(&bb, &conf.Conf.Common)
	if err != nil {
		log.Fatalf("Failed to execute %v: %v", template_html, err)
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
