package http

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/kiranraj27/sprint-planner/internal/domain"
)

type CardHandler struct {
	cardUseCase domain.CardUseCase
}

func NewCardHandler(router *mux.Router, uc domain.CardUseCase) {
	handler := &CardHandler{uc}
	router.HandleFunc("/cards", handler.GetCards).Methods("GET")
}

func (h *CardHandler) GetCards(w http.ResponseWriter, r *http.Request) {
	cards, err := h.cardUseCase.GetCards()
	if err != nil {
		http.Error(w, "Failed to fetch cards", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(cards)
}
