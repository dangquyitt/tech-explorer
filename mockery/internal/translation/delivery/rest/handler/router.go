package handler

import (
	"database/sql"
	"net/http"

	"github.com/dangquyitt/tech-explorer/mockery/internal/translation/infrastructure/postgresql"
	"github.com/dangquyitt/tech-explorer/mockery/internal/translation/infrastructure/webapi"
	"github.com/dangquyitt/tech-explorer/mockery/internal/translation/usecase"
)

func NewTranslationRouter(mux *http.ServeMux, conn *sql.Conn) {
	translationRepository := postgresql.NewTranslationRepository(conn)
	translationGoogle := webapi.NewTranslationGoogle()
	translationUsecase := usecase.NewTranslationUsecase(translationRepository, translationGoogle)
	translationHandler := &translationHandler{
		usecase: translationUsecase,
	}
	mux.HandleFunc("GET /translate", translationHandler.translate)
}
