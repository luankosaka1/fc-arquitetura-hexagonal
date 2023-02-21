package server

import (
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"github.com/luankosaka1/arquitetura-hexagonal-golang/adapters/web/handler"
	"github.com/luankosaka1/arquitetura-hexagonal-golang/application"
	"log"
	"net/http"
	"os"
	"time"
)

type Webserver struct {
	Service application.ProductServiceInterface
}

func MakeNewWebserver() *Webserver {
	return &Webserver{}
}

func (w Webserver) Server() {
	routers := mux.NewRouter()
	middleware := negroni.New(negroni.NewLogger())

	handler.MakeProductHandlers(routers, middleware, w.Service)
	http.Handle("/", routers)

	server := &http.Server{
		ReadHeaderTimeout: 10 * time.Second,
		WriteTimeout:      10 * time.Second,
		Addr:              ":9000",
		Handler:           http.DefaultServeMux,
		ErrorLog:          log.New(os.Stderr, "log", log.Lshortfile),
	}
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
