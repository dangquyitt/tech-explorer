package usecase

import (
	"context"
	"errors"
	"testing"

	"github.com/dangquyitt/tech-explorer/mockery/internal/translation/domain/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type TranslationUsecaseTestSuite struct {
	suite.Suite
	orgText                   string
	source                    string
	dest                      string
	ctx                       context.Context
	translationRepositoryMock *mocks.TranslationRepositoryMock
	translationWebAPIMock     *mocks.TranslationWebAPIMock
	translationUsecase        *translationUsecase
}

func (suite *TranslationUsecaseTestSuite) SetupTest() {
	suite.orgText = "original text"
	suite.source = "source"
	suite.dest = "dest"
	suite.ctx = context.Background()

	suite.translationRepositoryMock = new(mocks.TranslationRepositoryMock)
	suite.translationWebAPIMock = new(mocks.TranslationWebAPIMock)
	suite.translationUsecase = &translationUsecase{
		repository:        suite.translationRepositoryMock,
		translationWebAPI: suite.translationWebAPIMock,
	}
}

func TestTranslationUsecaseTestSuite(t *testing.T) {
	suite.Run(t, new(TranslationUsecaseTestSuite))
}

func (suite *TranslationUsecaseTestSuite) Test_Translate_WebAPIError_Error() {
	translateMock := suite.translationWebAPIMock.On("Translate", suite.ctx, suite.orgText, suite.source, suite.dest).Return(nil, errors.New(""))

	translation, err := suite.translationUsecase.Translate(suite.ctx, suite.orgText, suite.source, suite.dest)
	assert.Nil(suite.T(), translation)
	assert.NotNil(suite.T(), err)

	translateMock.Unset()
}
