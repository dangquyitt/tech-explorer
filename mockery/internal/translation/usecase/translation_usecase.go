package usecase

import (
	"context"
	"fmt"

	"github.com/dangquyitt/tech-explorer/mockery/internal/translation/domain"
)

type translationUsecase struct {
	repository        domain.TranslationRepository
	translationWebAPI domain.TranslationWebAPI
}

func (uc *translationUsecase) FetchHistories(ctx context.Context) ([]domain.Translation, error) {
	translations, err := uc.repository.FindHistories(ctx)
	if err != nil {
		return nil, fmt.Errorf("TranslationUseCase - History - s.repo.GetHistory: %w", err)
	}

	return translations, nil
}

func (uc *translationUsecase) Translate(ctx context.Context, orgText string, source string, dest string) (*domain.Translation, error) {
	translation, err := uc.translationWebAPI.Translate(ctx, orgText, source, dest)
	if err != nil {
		return nil, fmt.Errorf("TranslationUseCase - Translate - uc.translationWebAPI.Translate: %w", err)
	}

	err = uc.repository.Create(context.Background(), translation)
	if err != nil {
		return nil, fmt.Errorf("TranslationUseCase - Translate - s.repo.Store: %w", err)
	}

	return translation, nil
}

func NewTranslationUsecase(repository domain.TranslationRepository, translationWebAPI domain.TranslationWebAPI) domain.TranslationUsecase {
	return &translationUsecase{
		repository:        repository,
		translationWebAPI: translationWebAPI,
	}
}
