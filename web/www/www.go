package www

import (
	"bytes"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"

	_template "github.com/rasa/shortme/template"
)

const (
	html          = "/index.html"
	template_html = "template" + html
)

var bb bytes.Buffer

func Init() {
	bb.Reset()
	_template.Init()
	fh, err := _template.Assets.Open(html)
	if err != nil {
		log.Fatalf("Failed to open %v: %v", template_html, err)
	}
	defer fh.Close()
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

	err = tpl.Execute(&bb, &_template.Vars)
	if err != nil {
		log.Fatalf("Failed to execute %v: %v", template_html, err)
	}
	if bb.Len() == 0 {
		log.Fatalf("Failed to execute %v: %v", template_html, err)
	}
}

func Index(w http.ResponseWriter, _ *http.Request) {
	w.Write(bb.Bytes())
}
