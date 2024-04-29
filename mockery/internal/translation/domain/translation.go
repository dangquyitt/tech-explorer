package domain

import "context"

type Translation struct {
	Source      string `json:"source"`
	Destination string `json:"destination"`
	Original    string `json:"original"`
	Translation string `json:"translation"`
}

type TranslationUsecase interface {
	Translate(ctx context.Context, orgText, source, dest string) (*Translation, error)
	FetchHistories(ctx context.Context) ([]Translation, error)
}

type TranslationRepository interface {
	Create(ctx context.Context, translation *Translation) error
	FindHistories(ctx context.Context) ([]Translation, error)
}

type TranslationWebAPI interface {
	Translate(ctx context.Context, orgText, source, dest string) (*Translation, error)
}
