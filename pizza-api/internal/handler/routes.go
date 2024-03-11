package handler

import (
	
	"github.com/gorilla/mux"
)

func (h *Handler) InitRoutes() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/orders", h.OrderHandler)
	r.HandleFunc("/orders/exec", h.OrderExecHandler)

	r.HandleFunc("/auth/signUp", h.SignUpHandler)
	r.HandleFunc("/auth/signIn", h.SignInHandler)

	return r
}
