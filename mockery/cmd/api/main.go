package main

import (
	"log"
	"net/http"

	"github.com/dangquyitt/tech-explorer/mockery/internal/translation/delivery/rest/handler"
)

func main() {
	mux := http.NewServeMux()
	handler.NewTranslationRouter(mux, nil)
	log.Fatal(http.ListenAndServe(":8080", mux))
}
