package user

import "github.com/gorilla/mux"

func RegisterRoutes(r *mux.Router, h *Handler) {
	r.HandleFunc("/registration", h.Registration).Methods("POST")
	// r.HandleFunc("/login", h.Login).Methods("POST")
	// r.HandleFunc("/logout", h.Logout).Methods("POST")
}
