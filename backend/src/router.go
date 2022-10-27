package main

import (
	"net/http"
)

type Router struct {
	handler *Handler
}

func NewRouter(repository *Repository) *Router {
	return &Router{
		handler: NewHandler(repository),
	}
}

func (I *Router) init() {
	http.HandleFunc("/", I.handler.Home)
	http.HandleFunc("/locker", I.handler.ChangeState)
}
