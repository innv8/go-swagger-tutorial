package controllers

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

type Base struct {
	Router *mux.Router
}

func (b *Base) Init() {
	b.Router = mux.NewRouter()

	b.Router.HandleFunc("/persons", b.GetPersons).Methods("GET")
	b.Router.HandleFunc("/persons", b.AddPerson).Methods("POST")

	b.Router.PathPrefix("/docs/").Handler(b.DocsHandler()).Methods("GET")

	var host = "0.0.0.0:5000"
	log.Printf("starting API on %s", host)
	var c = cors.New(cors.Options{
		AllowCredentials: true,
		AllowedOrigins:   []string{"*"},
	})
	var handler = c.Handler(b.Router)
	log.Panicln(http.ListenAndServe(host, handler))

}
