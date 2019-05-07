package controller

import (
	"github.com/gorilla/mux"
	"sync"
)

type Controller struct {
	root string

	router *mux.Router
	once   sync.Once
}

func New(root string) *Controller {
	return &Controller{
		root: root,
	}
}

func (c *Controller) Router() *mux.Router {
	c.once.Do(func() {
		r := mux.NewRouter()

		r.HandleFunc("/", c.RootHandler)
		r.HandleFunc("/do", c.DoHandler).Methods("POST")

		c.router = r
	})
	return c.router
}
