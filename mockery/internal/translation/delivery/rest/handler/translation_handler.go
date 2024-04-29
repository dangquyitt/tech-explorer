package handler

import (
	"net/http"

	"github.com/dangquyitt/tech-explorer/mockery/internal/translation/domain"
)

type translationHandler struct {
	usecase domain.TranslationUseCase
}

type translateRequest struct {
	Source      string `json:"source"`
	Destination string `json:"destination"`
	Original    string `json:"original"`
}

func (t *translationHandler) translate(w http.ResponseWriter, r *http.Request) {
	var rq translateRequest

	translation, err := t.usecase.Translate(r.Context(), rq.Original, rq.Source, rq.Destination)
	if err != nil {
		http.Error(w, "Failed to decode request body", http.StatusBadRequest)
		return
	}

	renderJSON(w, translation)
}
