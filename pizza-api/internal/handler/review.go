package handler

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"time"
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

	review := struct {
		UserId int    `json:"userid"`
		Text   string `json:"text"`
	}{}

	err = json.Unmarshal(data, &review)
	if err != nil {
		h.log.SugaredLogger.Fatalf("unmarshal json error: %w", err)
		w.WriteHeader(http.StatusInternalServerError)
	}

	err = h.service.APIProvider.CreateReview(ctx, review.UserId, review.Text)
	if err != nil {
		h.log.SugaredLogger.Fatalf("create review error: %w", err)
		w.WriteHeader(http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusOK)
}
