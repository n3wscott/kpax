package controller

import (
	"html/template"
	"net/http"
)

func (c *Controller) RootHandler(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles(c.root + "/templates/index.html")
	_ = t.Execute(w, nil)
}
