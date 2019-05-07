package controllers

import (
	"html/template"
	"net/http"
)

func RootHandler(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("/home/nicholss/go/src/github.com/n3wscott/kpax/poke/cmd/poke/kodata/templates/index.html")
	t.Execute(w, nil)
}
