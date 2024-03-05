package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	grpcapp "github.com/eeQuillibrium/pizza-api/internal/app/grpc"
	"github.com/eeQuillibrium/pizza-api/internal/domain/models"
	nikita_auth1 "github.com/eeQuillibrium/protos/gen/go/auth"
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

	r.HandleFunc("/orders", h.OrderHandler)
	r.HandleFunc("/orders/exec", h.OrderExecHandler)

	r.HandleFunc("/auth/signUp", h.SignUpHandler)
	r.HandleFunc("/auth/signIn", h.SignInHandler)

	return r
}

func (h *Handler) OrderHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Order request, method:", r.Method+"...")
}

func (h *Handler) OrderExecHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("OrderExec request, method:", r.Method+"...")

	var order models.Order
	b, err := io.ReadAll(r.Body)
	//err := json.NewDecoder(r.Body).Decode(&order)
	if err != nil {
		log.Fatalf("order json decode problem: %v", err)
	}
	log.Print(string(b))
	json.Unmarshal(b, &order)

	log.Printf("order parameters: UserId: %d Price: %0.2f Units: %v", order.UserId, order.Price, order.Units)

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

	token, err := h.GRPCapp.Auth.Login(ctx, &nikita_auth1.LoginRequest{Login: u.Login, Pass: u.Password, AppId: 1})
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
