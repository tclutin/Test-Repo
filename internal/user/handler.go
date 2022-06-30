package user

import (
	"REST/internal/handlers"
	"net/http"
)

type handler struct{}

func NewHandler() handlers.Handler {
	return &handler{}
}

func (h *handler) Register(router *http.ServeMux) {
	router.HandleFunc("/login", h.APILogin)
	router.HandleFunc("/register", h.APIRegister)
	router.HandleFunc("/logout", h.APILogout)
}

func (h *handler) APILogin(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	w.Write([]byte("login"))
}

func (h *handler) APIRegister(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	w.Write([]byte("Register"))
}

func (h *handler) APILogout(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	w.Write([]byte("Logout"))
}
