package handler

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	nikita_auth1 "github.com/eeQuillibrium/protos/gen/go/auth"

	"github.com/eeQuillibrium/pizza-api/internal/domain/models"
)

func (h *Handler) authHandler(w http.ResponseWriter, r *http.Request) {
}
func (h *Handler) signUpHandler(w http.ResponseWriter, r *http.Request) {
	h.log.SugaredLogger.Info("request for signUp...")

	u := models.User{}

	err := json.NewDecoder(r.Body).Decode(&u)
	if err != nil {
		h.log.SugaredLogger.Fatalf("decode json problem %w", err)
	}

	h.log.SugaredLogger.Infof("successful user shared data phone: %s, pass: %s", u.Phone, u.Password)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	userId, err := h.GRPCApp.Auth.Register(ctx, &nikita_auth1.RegRequest{Phone: u.Phone, Pass: u.Password})
	if err != nil {
		h.log.SugaredLogger.Fatalf("registration error: %w", err)
	}

	u.UserId = int(userId)

	if err := h.service.APIProvider.StoreUser(ctx, &u); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		h.log.SugaredLogger.Fatalf("gateway user storing error: %w", err)
	}

	h.log.SugaredLogger.Infof("user with id=%d was registered completely!", userId)
	http.Redirect(w, r, "http://localhost:82/home", http.StatusPermanentRedirect)
}

func (h *Handler) signInHandler(w http.ResponseWriter, r *http.Request) {
	h.log.SugaredLogger.Info("request for signIn...")

	var u models.User

	err := json.NewDecoder(r.Body).Decode(&u)
	if err != nil {
		h.log.SugaredLogger.Fatal(err)
	}

	h.log.SugaredLogger.Infof("successful user shared data login: %s, pass: %s", u.Phone, u.Password)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	token, err := h.GRPCApp.Auth.Login(ctx, &nikita_auth1.LoginRequest{Phone: u.Phone, Pass: u.Password, AppId: 1})
	if err != nil {
		h.log.SugaredLogger.Fatalf("login error: %w", err)
	}

	h.log.SugaredLogger.Infof("successful login! token: %s", token)

	http.SetCookie(w, &http.Cookie{
		Name:  "token_bearer",
		Value: token,
	})

	http.Redirect(w, r, "localhost", http.StatusPermanentRedirect)
}
