package api

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"

	"github.com/yougg/shortme/conf"
)

func CheckVersion(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	versionInfo, _ := json.Marshal(version{Version: conf.Version})
	w.Write(versionInfo)
}

func CheckHealth(w http.ResponseWriter, r *http.Request) {
	tpl := template.New("health.html")
	var err error
	tpl, err = tpl.ParseFiles("template/health.html")
	if err != nil {
		log.Printf("parse template error. %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(http.StatusText(http.StatusInternalServerError)))
		return
	}
	err = tpl.Execute(w, nil)
	if err != nil {
		log.Printf("execute template error. %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(http.StatusText(http.StatusInternalServerError)))
	}
}
