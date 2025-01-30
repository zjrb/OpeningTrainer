package handler

import (
	"encoding/json"
	"net/http"

	"github.com/zjrb/OpeningTrainer/internal/core/services"
	"github.com/zjrb/OpeningTrainer/internal/logger"
)

type OpeningHandler struct {
	svc    *services.OpeningService
	logger *logger.Logger
}

func NewOpeningHandler(svc *services.OpeningService) *OpeningHandler {
	return &OpeningHandler{svc: svc}
}

func (o *OpeningHandler) GetOpening() http.HandlerFunc {
	type openingRes struct {
		OpeningName string `json:"opening_name"`
		ECO         string `json:"eco"`
		PGN         string `json:"pgn"`
		UCI         string `json:"uci"`
		FEN         string `json:"fen"`
	}
	type response struct {
		Openings []openingRes `json:"openings"`
	}
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		name := r.PathValue("name")
		if name == "" {
			http.Error(w, "Missing 'name' query parameter", http.StatusBadRequest)
			return
		}
		openings := o.svc.GetOpeningByName(name)
		if len(openings) == 0 {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(response{Openings: []openingRes{}})
			return
		}
		resp := response{Openings: make([]openingRes, len(openings))}
		for i, opening := range openings {
			resp.Openings[i] = openingRes{
				OpeningName: opening.OpeningName,
				ECO:         opening.ECO,
				PGN:         opening.PGN,
				UCI:         opening.UCI,
				FEN:         opening.FEN,
			}
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(resp)
	})
}
