package controller

import (
	"fmt"
	"mime"
	"net/http"
	"time"
)

func (c *Controller) DoHandler(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case "POST":

		ct := r.Header.Get("Content-Type")
		ct, _, err := mime.ParseMediaType(ct)
		if err != nil {
			fmt.Fprintf(w, "ParseMediaType() err: %v", err)
		}

		switch {
		case ct == "multipart/form-data":
			if err := r.ParseMultipartForm(1000); err != nil {
				fmt.Fprintf(w, "ParseMultipartForm() err: %v", err)
				return
			}
		default:
			fmt.Fprintf(w, "unhandled: %v", ct)
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

		fmt.Fprintf(w, "Time = %s\n", time.Now().String())
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
