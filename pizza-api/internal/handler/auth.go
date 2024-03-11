package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	nikita_auth1 "github.com/eeQuillibrium/protos/gen/go/auth"

	"github.com/eeQuillibrium/pizza-api/internal/domain/models"
)

func (h *Handler) SignUpHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("request for signUp...")

	var u models.User

	err := json.NewDecoder(r.Body).Decode(&u)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("successful user shared data login: %v, pass: %v", u.Login, u.Password)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	userId, err := h.GRPCApp.Auth.Register(ctx, &nikita_auth1.RegRequest{Login: u.Login, Pass: u.Password})
	if err != nil {
		log.Fatalf("registration error: %v", err)
	}

	log.Printf("user with id=%d was registered completely!", userId)

	w.Write([]byte(fmt.Sprintf("user with id=%d was registered completely!", userId)))
}

func (h *Handler) SignInHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("request for signIn...")

	var u models.User

	err := json.NewDecoder(r.Body).Decode(&u)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("successful user shared data login: %v, pass: %v", u.Login, u.Password)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	token, err := h.GRPCApp.Auth.Login(ctx, &nikita_auth1.LoginRequest{Login: u.Login, Pass: u.Password, AppId: 1})
	if err != nil {
		log.Fatalf("login error: %v", err)
	}

	log.Printf("successful login! token: %s", token)

	http.SetCookie(w, &http.Cookie{
		Name:  "token_bearer",
		Value: token,
	})

	http.Redirect(w, r, "localhost", http.StatusPermanentRedirect)
}
