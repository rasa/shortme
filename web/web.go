package web

import (
	"fmt"
	"log"
	"net/http"
	"regexp"

	"github.com/gorilla/mux"
	"github.com/rasa/shortme/conf"
	"github.com/rasa/shortme/static"
	"github.com/rasa/shortme/web/api"
	"github.com/rasa/shortme/web/www"
	_www "github.com/rasa/shortme/www"
)

func Start() {
	log.Println("web starts")
 
	api.Init()
	www.Init()
	r := mux.NewRouter()

	r.HandleFunc("/version", api.CheckVersion).Methods(http.MethodGet)
	r.HandleFunc("/health", api.CheckHealth).Methods(http.MethodGet)
	r.HandleFunc("/short", api.ShortURL).Methods(http.MethodPost).HeadersRegexp("Content-Type", "application/json")
	r.HandleFunc("/expand", api.ExpandURL).Methods(http.MethodPost).HeadersRegexp("Content-Type", "application/json")

	r.HandleFunc("/", www.Index).Methods(http.MethodGet)
	r.HandleFunc("/index.html", www.Index).Methods(http.MethodGet)

	r.Handle("/static/{type}/{file}", http.StripPrefix("/static/", http.FileServer(static.Assets)))
	r.Handle("/favicon.ico", http.StripPrefix("/", http.FileServer(_www.Assets)))
	r.Handle("/robots.txt", http.StripPrefix("/", http.FileServer(_www.Assets)))

	shortenedURL := fmt.Sprintf("/{shortenedURL:[%v]{1,%v}}", regexp.QuoteMeta(conf.Conf.Common.BaseString), conf.Conf.Common.ShortURLMax)

	r.HandleFunc(shortenedURL, api.Redirect).Methods(http.MethodGet)

	log.Fatal(http.ListenAndServe(conf.Conf.Http.Listen, r))
}
