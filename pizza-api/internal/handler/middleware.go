package handler

import (
	"context"
	"fmt"
	"net/http"

	nikita_auth1 "github.com/eeQuillibrium/protos/gen/go/auth"
)

func (h *Handler) userIdentify(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("token_bearer")
		if token == "" {
			http.Error(w, "empty token", http.StatusUnauthorized)
		}

		userId, err := h.GRPCApp.Auth.UserIdentify(
			context.Background(),
			&nikita_auth1.IdentifyRequest{Token: token},
		)
		if err != nil {
			h.log.SugaredLogger.Infof("error from auth: %v", err)
			http.Error(w, "error with identify", http.StatusUnauthorized)
		}

		r.Header.Set("userId", fmt.Sprintf("%d", userId))

		next.ServeHTTP(w, r)
	})
}
