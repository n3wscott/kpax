package controllers

import (
	"encoding/json"
	"net/http"
)

func DoHandler(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	request := make(map[string]interface{}, 0)
	err := decoder.Decode(&request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	return
}
