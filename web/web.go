package web

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"regexp"
	"sync/atomic"
	"syscall"
	"time"

	"github.com/gorilla/mux"
	"github.com/rasa/shortme/conf"
	"github.com/rasa/shortme/static"
	"github.com/rasa/shortme/web/api"
	"github.com/rasa/shortme/web/www"
	_www "github.com/rasa/shortme/www"
)

// inspired by https://stackoverflow.com/a/43639645

type myServer struct {
	http.Server
	shutdownReq chan bool
	reqCount    uint32
	sighupped   uint32
}

func NewServer() *myServer {
	//create server
	s := &myServer{
		Server: http.Server{
			Addr:         conf.Conf.Http.Listen,
			ReadTimeout:  10 * time.Second,
			WriteTimeout: 10 * time.Second,
		},
		shutdownReq: make(chan bool),
	}

	//register handlers
	r := mux.NewRouter()

	r.HandleFunc("/expand", api.ExpandURL).Methods(http.MethodPost).HeadersRegexp("Content-Type", "application/json")
	r.HandleFunc("/health", api.CheckHealth).Methods(http.MethodGet)
	r.HandleFunc("/short", api.ShortURL).Methods(http.MethodPost).HeadersRegexp("Content-Type", "application/json")
	r.HandleFunc("/version", api.CheckVersion).Methods(http.MethodGet)
	r.HandleFunc("/shutdown", s.ShutdownHandler)

	r.HandleFunc("/", www.Index).Methods(http.MethodGet)
	r.HandleFunc("/index.html", www.Index).Methods(http.MethodGet)

	r.Handle("/static/{type}/{file}", http.StripPrefix("/static/", http.FileServer(static.Assets)))
	r.Handle("/favicon.ico", http.StripPrefix("/", http.FileServer(_www.Assets)))
	r.Handle("/robots.txt", http.StripPrefix("/", http.FileServer(_www.Assets)))

	shortenedURL := fmt.Sprintf("/{shortenedURL:[%v]{1,%v}}", regexp.QuoteMeta(conf.Conf.Common.BaseString), conf.Conf.Common.ShortURLMax)

	r.HandleFunc(shortenedURL, api.Redirect).Methods(http.MethodGet)

	//set http server handler
	s.Handler = r

	return s
}

func (s *myServer) WaitShutdown() {
	irqSig := make(chan os.Signal, 1)
	signal.Notify(irqSig,
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGTERM,
	)

	// log.Printf("Waiting for shutdown signal")
	//Wait interrupt or shutdown request through /shutdown
	select {
	case sig := <-irqSig:
		log.Printf("Received shutdown request: %v", sig)
		atomic.CompareAndSwapUint32(&s.reqCount, 0, 1)
		if sig == syscall.SIGHUP {
			atomic.CompareAndSwapUint32(&s.sighupped, 0, 1)
		}
	case sig := <-s.shutdownReq:
		log.Printf("Shutdown shutdown request via /shutdown URL: %v", sig)
	}

	log.Printf("Stoping http server")

	//Create shutdown context with 10 second timeout
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	//shutdown the server
	err := s.Shutdown(ctx)
	if err != nil {
		log.Printf("Shutdown request error: %v", err)
	}
}

func (s *myServer) ShutdownHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Shutting down http server"))

	//Do nothing if shutdown request already issued
	//if s.reqCount == 0 then set to 1, return true otherwise false
	if !atomic.CompareAndSwapUint32(&s.reqCount, 0, 1) {
		log.Printf("Shutdown via API call already in progress")
		return
	}

	go func() {
		s.shutdownReq <- true
	}()
}

func (s *myServer) Running() bool {
	return atomic.LoadUint32(&s.reqCount) == 0
}

func (s *myServer) Sighupped() bool {
	return atomic.LoadUint32(&s.sighupped) == 1
}

func Start() bool {
	log.Println("web starts")

	api.Init()
	www.Init()

	//Start the server
	server := NewServer()

	done := make(chan bool)
	go func() {
		log.Printf("Starting http server on %v", conf.Conf.Http.Listen)
		err := server.ListenAndServe()
		if err != nil {
			if server.Running() {
				log.Printf("ListenAndServe() failed: %v", err)
			}
		}
		done <- true
	}()

	//wait shutdown
	server.WaitShutdown()
	log.Printf("Web server shutdown, gracefully exiting")
	return server.Sighupped()
}
