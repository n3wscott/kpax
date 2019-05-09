package controller

import (
	"html/template"
	"net/http"
	"sync"
)

var once sync.Once
var t *template.Template

func (c *Controller) RootHandler(w http.ResponseWriter, r *http.Request) {

	once.Do(func() {
		t, _ = template.ParseFiles(
			c.root+"/templates/index.html",
			c.root+"/templates/top.html",
			c.root+"/templates/side.html",
		)
	})

	data := map[string]interface{}{
		"test": "data",
	}

	eventTypes := c.kn.EventTypes(c.namespace)

	data["eventTypes"] = eventTypes

	_ = t.Execute(w, data)
}

func getQueryParam(r *http.Request, key string) string {
	keys, ok := r.URL.Query()[key]
	if !ok || len(keys[0]) < 1 {
		return ""
	}
	return keys[0]
}
