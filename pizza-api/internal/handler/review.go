package handler

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"time"

	"github.com/eeQuillibrium/pizza-api/internal/domain/models"
)

func (h *Handler) reviewHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("reviewPage")) //send template
}

func (h *Handler) reviewSendHandler(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*4)
	defer cancel()

	data, err := io.ReadAll(r.Body)
	if err != nil {
		h.log.SugaredLogger.Fatalf("read body error: %w", err)
		w.WriteHeader(http.StatusInternalServerError)
	}

	review := models.Review{}

	err = json.Unmarshal(data, &review)
	if err != nil {
		h.log.SugaredLogger.Fatalf("unmarshal json error: %w", err)
		w.WriteHeader(http.StatusInternalServerError)
	}

	err = h.service.APIProvider.CreateReview(ctx, &review)
	if err != nil {
		h.log.SugaredLogger.Fatalf("create review error: %w", err)
		w.WriteHeader(http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusOK)
}
