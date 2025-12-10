package client

import "github.com/gorilla/mux"

func RegisterRoutes(r *mux.Router, h *Handler) {
	r.HandleFunc("/clients", h.Get).Methods("GET")
	r.HandleFunc("/clients", h.Create).Methods("POST")
	r.HandleFunc("/clients/{id}", h.Update).Methods("PATCH")
	r.HandleFunc("/clients/{id}", h.Delete).Methods("DELETE")
}
