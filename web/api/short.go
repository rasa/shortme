package api

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"

	"github.com/gorilla/mux"
	"github.com/rasa/shortme/conf"
	"github.com/rasa/shortme/short"
)

func Redirect(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	shortededURL := vars["shortenedURL"]

	longURL, err := short.Shorter.Expand(shortededURL)
	if err != nil {
		log.Printf("Failed to redirect %v: %v", shortededURL, err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(http.StatusText(http.StatusInternalServerError)))
		return
	}

	if len(longURL) != 0 {
		w.Header().Set("Location", longURL)
		w.WriteHeader(http.StatusTemporaryRedirect)
	} else {
		w.WriteHeader(http.StatusNoContent)
	}
}

func ShortURL(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf("Read short request error: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		errMsg, _ := json.Marshal(errorResp{Msg: http.StatusText(http.StatusInternalServerError)})
		w.Write(errMsg)
		return
	}

	var shortReq shortReq
	err = json.Unmarshal(body, &shortReq)
	if err != nil {
		log.Printf("Parse short request error: %v", err)
		w.WriteHeader(http.StatusBadRequest)
		errMsg, _ := json.Marshal(errorResp{Msg: http.StatusText(http.StatusBadRequest)})
		w.Write(errMsg)
		return
	}

	//var longURL *url.URL
	shortReq.LongURL = strings.TrimSpace(shortReq.LongURL)
	
	longURL, err := url.Parse(shortReq.LongURL)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		errMsg, _ := json.Marshal(errorResp{Msg: "requested url is malformed"})
		w.Write(errMsg)
		return
	}

	if longURL.Scheme == "" {
		longURL.Scheme = "http"
	}
	
	if longURL.Host == "" {
		w.WriteHeader(http.StatusBadRequest)
		errMsg, _ := json.Marshal(errorResp{Msg: "requested url is malformed"})
		w.Write(errMsg)
		return
	}
	
	longURL.Scheme := strings.ToLower(longURL.Scheme)
	longURL.Host := strings.ToLower(longURL.Host)
	
	if longURL.Host == strings.ToLower(conf.Conf.Common.DomainName) {
		w.WriteHeader(http.StatusBadRequest)
		errMsg, _ := json.Marshal(errorResp{Msg: "requested url is already shortened"})
		w.Write(errMsg)
		return
	}
	
	if longURL.Scheme != "http" && longURL.Scheme != "https" {
		w.WriteHeader(http.StatusBadRequest)
		errMsg, _ := json.Marshal(errorResp{Msg: "requested url is not a http or https url"})
		w.Write(errMsg)
		return
	}

	var shortenedURL string
	shortenedURL, err = short.Shorter.Short(longURL.String())
	if err != nil {
		log.Printf("Short url error: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		errMsg, _ := json.Marshal(errorResp{Msg: http.StatusText(http.StatusInternalServerError)})
		w.Write(errMsg)
		return
	}
	shortenedURL = (&url.URL{
		Scheme: conf.Conf.Common.Schema,
		Host:   conf.Conf.Common.DomainName,
		Path:   shortenedURL,
	}).String()
	shortResp, _ := json.Marshal(shortResp{ShortURL: shortenedURL})
	w.Write(shortResp)
}

func ExpandURL(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf("Read expand request error: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		errMsg, _ := json.Marshal(errorResp{Msg: http.StatusText(http.StatusInternalServerError)})
		w.Write(errMsg)
		return
	}

	var expandReq expandReq
	err = json.Unmarshal(body, &expandReq)
	if err != nil {
		log.Printf("Parse expand request error: %v", err)
		w.WriteHeader(http.StatusBadRequest)
		errMsg, _ := json.Marshal(errorResp{Msg: http.StatusText(http.StatusBadRequest)})
		w.Write(errMsg)
		return
	}

	var shortURL *url.URL
	shortURL, err = url.Parse(expandReq.ShortURL)
	if err != nil {
		log.Printf(`Invalid URL: "%v"`, expandReq.ShortURL)
		w.WriteHeader(http.StatusBadRequest)
		errMsg, _ := json.Marshal(errorResp{Msg: http.StatusText(http.StatusBadRequest)})
		w.Write(errMsg)
		return
	}

	var expandedURL string
	expandedURL, err = short.Shorter.Expand(strings.TrimLeft(shortURL.Path, "/"))
	if err != nil {
		log.Printf("Failed to expand %v: %v", shortURL.Path, err)
		w.WriteHeader(http.StatusInternalServerError)
		errMsg, _ := json.Marshal(errorResp{Msg: http.StatusText(http.StatusInternalServerError)})
		w.Write(errMsg)
		return
	}

	expandResp, _ := json.Marshal(expandResp{LongURL: expandedURL})
	w.Write(expandResp)
}
