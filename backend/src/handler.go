package main

import (
	"net/http"

	utilHttp "github.com/pedroh-reis/valu-safe/backend/src/http"
)

type Handler struct {
	resolver Resolver
}

func NewHandler(repository *Repository) *Handler {
	return &Handler{
		resolver: *NewResolver(repository),
	}
}

func (I *Handler) Home(w http.ResponseWriter, r *http.Request) {
	utilHttp.ExecWithoutInput(w, r, I.resolver.Home)
}

func (I *Handler) ChangeState(w http.ResponseWriter, r *http.Request) {
	utilHttp.Exec(w, r, I.resolver.ChangeState)
}

func (I *Handler) GetState(w http.ResponseWriter, r *http.Request) {
	utilHttp.Exec(w, r, I.resolver.GetState)
}

func (I *Handler) GetStatistics(w http.ResponseWriter, r *http.Request) {
	utilHttp.Exec(w, r, I.resolver.GetStatistics)
}
