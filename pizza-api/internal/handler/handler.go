package handler

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	grpcapp "github.com/eeQuillibrium/pizza-api/internal/app/grpc"
	"github.com/eeQuillibrium/pizza-api/internal/domain/models"
	nikita_auth1 "github.com/eeQuillibrium/protos/proto/gen/go/auth"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

type Handler struct {
	GRPCapp *grpcapp.GRPCApp
}

func New(
	port string,
) *Handler {
	grpcapp := grpcapp.New(port)
	return &Handler{grpcapp}
}

func (h *Handler) InitRoutes() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/", h.HomeHandler)
	r.HandleFunc("/contacts", h.ContactsHandler)
	r.HandleFunc("/auth/signUp", h.SignUpHandler)
	r.HandleFunc("/auth/signIn", h.SignInHandler)
	return r
}

func (h *Handler) HomeHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Homepage request, method:", r.Method+"...")

}
func (h *Handler) ContactsHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("contactpage")
}
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

	userId, err := h.GRPCapp.Auth.Register(ctx, &nikita_auth1.RegRequest{Login: u.Login, Pass: u.Password})
	if err != nil { 
		log.Fatalf("registration error: %v", err)
	}

	log.Print("user with id %v was registered completely", userId)
}
func (h *Handler) SignInHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("")
}
