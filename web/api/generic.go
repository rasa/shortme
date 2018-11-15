package api

import (
	"bytes"
	"encoding/json"
	"html/template"
	"log"
	"net/http"

	"github.com/rasa/shortme/conf"
)

const (
	html          = "health.html"
	template_html = "template" + html
)

var bb bytes.Buffer

func Init() {
	tpl := template.New(html)
	var err error
	tpl, err = tpl.ParseFiles(template_html)
	if err != nil {
		log.Fatalf("Failed to parse %v: %v", template_html, err)
	}

	err = tpl.Execute(&bb, &conf.Conf.Common)
	if err != nil {
		log.Fatalf("Failed to execute %v: %v", template_html, err)
	}
}

func CheckVersion(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	versionInfo, _ := json.Marshal(version{Version: conf.Version})
	w.Write(versionInfo)
}

func CheckHealth(w http.ResponseWriter, _ *http.Request) {
	if bb.Len() == 0 {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(http.StatusText(http.StatusInternalServerError)))
		return
	}
	w.Write(bb.Bytes())
}
