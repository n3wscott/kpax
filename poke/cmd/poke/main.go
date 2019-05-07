package main

import (
	"log"
	"net/http"

	"github.com/n3wscott/kpax/poke/pkg/controller"
)

func main() {

	port := ":8080"
	root := "/Users/nicholss/go/src/github.com/n3wscott/kpax/poke/cmd/poke/kodata"

	c := controller.New(root)

	c.Router().PathPrefix("/static/").Handler(http.StripPrefix("/static/",
		http.FileServer(http.Dir(root+"/static"))))

	http.Handle("/", c.Router())
	log.Fatal(http.ListenAndServe(port, nil))
}
