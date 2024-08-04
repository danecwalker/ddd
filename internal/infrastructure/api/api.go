package api

import (
	"log"
	"net/http"
)

type ApiHandler struct {
	Mux *http.ServeMux
}

func NewApiHandler() *ApiHandler {
	mux := http.NewServeMux()

	return &ApiHandler{
		Mux: mux,
	}
}

func (h *ApiHandler) Serve(address string) error {
	log.Printf("Listening on %s", address)
	return http.ListenAndServe(address, h.Mux)
}

func (h *ApiHandler) NewSubRouter(path string) *http.ServeMux {
	subRouter := http.NewServeMux()
	h.Mux.Handle(path, subRouter)

	return subRouter
}
