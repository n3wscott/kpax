package controller

import (
	"github.com/gorilla/mux"
	"github.com/n3wscott/knap/pkg/knative"
	"sync"
)

type Controller struct {
	root      string
	kn        *knative.Client
	namespace string

	router *mux.Router
	once   sync.Once
}

func New(root string, kn *knative.Client, ns string) *Controller {
	return &Controller{
		root:      root,
		kn:        kn,
		namespace: ns,
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
