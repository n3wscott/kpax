package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/n3wscott/kpax/poke/pkg/controllers"
)

func main() {

	port := ":8080"

	r := mux.NewRouter()
	r.HandleFunc("/", controllers.RootHandler)
	r.HandleFunc("/do", controllers.DoHandler).Methods("POST")
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/",
		http.FileServer(http.Dir("/home/nicholss/go/src/github.com/n3wscott/kpax/poke/cmd/poke/kodata/static"))))

	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(port, nil))
}
