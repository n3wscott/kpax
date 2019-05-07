package controller

import (
	"fmt"
	"net/http"
)

func (c *Controller) DoHandler(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case "POST":
		if err := r.ParseForm(); err != nil {
			fmt.Fprintf(w, "ParseForm() err: %v", err)
			return
		}
		fmt.Fprintf(w, "Post from website! r.PostFrom = %v\n", r.PostForm)
		eventtype := r.FormValue("eventtype")

		for k, v := range r.PostForm {
			fmt.Println(k, "~~>", v)
		}

		for k, v := range r.Form {
			fmt.Println(k, "-->", v)
		}

		fmt.Fprintf(w, "Form = %+v\n", r.PostForm)

		fmt.Fprintf(w, "Type = %s\n", eventtype)
	default:
		fmt.Fprintf(w, "only POST method is supported.")
	}

	//decoder := json.NewDecoder(r.Body)
	//request := make(map[string]interface{}, 0)
	//err := decoder.Decode(&request)
	//if err != nil {
	//	http.Error(w, err.Error(), http.StatusInternalServerError)
	//	return
	//}
	return
}
