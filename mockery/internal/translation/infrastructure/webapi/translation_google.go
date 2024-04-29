package webapi

import (
	"context"
	"fmt"

	translator "github.com/Conight/go-googletrans"
	"github.com/dangquyitt/tech-explorer/mockery/internal/translation/domain"
)

type translationGoogle struct {
	config translator.Config
}

func (t *translationGoogle) Translate(ctx context.Context, orgText string, source string, dest string) (*domain.Translation, error) {
	trans := translator.New(t.config)

	result, err := trans.Translate(orgText, source, dest)
	if err != nil {
		return nil, fmt.Errorf("TranslationWebAPI - Translate - trans.Translate: %w", err)
	}

	return &domain.Translation{
		Original:    orgText,
		Source:      source,
		Destination: dest,
		Translation: result.Text,
	}, nil
}

func NewTranslationGoogle() domain.TranslationWebAPI {
	return &translationGoogle{
		config: translator.Config{
			UserAgent:   []string{"Mozilla/5.0 (X11; Ubuntu; Linux x86_64; rv:15.0) Gecko/20100101 Firefox/15.0.1"},
			ServiceUrls: []string{"translate.google.com"},
		},
	}
}
