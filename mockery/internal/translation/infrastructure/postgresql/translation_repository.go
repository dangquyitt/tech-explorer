package postgresql

import (
	"context"
	"database/sql"

	"github.com/dangquyitt/tech-explorer/mockery/internal/translation/domain"
)

type translationRepository struct {
	conn *sql.Conn
}

func (t *translationRepository) Create(ctx context.Context, translation *domain.Translation) error {
	panic("unimplemented")
}

func (t *translationRepository) FindHistories(ctx context.Context) ([]domain.Translation, error) {
	panic("unimplemented")
}

func NewTranslationRepository(conn *sql.Conn) domain.TranslationRepository {
	return &translationRepository{
		conn: conn,
	}
}
