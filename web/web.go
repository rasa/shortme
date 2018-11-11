package web

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rasa/shortme/conf"
	"github.com/rasa/shortme/web/api"
	"github.com/rasa/shortme/web/www"
)

func Start() {
	log.Println("web starts")
	r := mux.NewRouter()

	r.HandleFunc("/version", api.CheckVersion).Methods(http.MethodGet)
	r.HandleFunc("/health", api.CheckHealth).Methods(http.MethodGet)
	r.HandleFunc("/short", api.ShortURL).Methods(http.MethodPost).HeadersRegexp("Content-Type", "application/json")
	r.HandleFunc("/expand", api.ExpandURL).Methods(http.MethodPost).HeadersRegexp("Content-Type", "application/json")
	r.HandleFunc("/{shortenedURL:[a-zA-Z0-9]{1,11}}", api.Redirect).Methods(http.MethodGet)

	r.HandleFunc("/", www.Index).Methods(http.MethodGet)
	r.HandleFunc("/index.html", www.Index).Methods(http.MethodGet)

	r.Handle("/static/{type}/{file}", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	r.Handle("/favicon.ico", http.StripPrefix("/", http.FileServer(http.Dir("."))))
	r.Handle("/robots.txt", http.StripPrefix("/", http.FileServer(http.Dir("."))))

	log.Fatal(http.ListenAndServe(conf.Conf.Http.Listen, r))
}
