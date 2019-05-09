package controller

import (
	"html/template"
	"net/http"
)

func (c *Controller) RootHandler(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles(c.root + "/templates/index.html")

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
