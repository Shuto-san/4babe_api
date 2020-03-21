package interactor

import (
	"github.com/Shuto-san/4babe_api/domain/model"
	"github.com/Shuto-san/4babe_api/usecase/presenter"
	"github.com/Shuto-san/4babe_api/usecase/repository"
)

type questionInteractor struct {
	QuestionRepository repository.QuestionRepository
	QuestionPresenter  presenter.QuestionPresenter
}

type QuestionInteractor interface {
	Get(u []*model.Question) ([]*model.Question, error)
}

func NewQuestionInteractor(r repository.QuestionRepository, p presenter.QuestionPresenter) QuestionInteractor {
	return &questionInteractor{r, p}
}

func (us *questionInteractor) Get(u []*model.Question) ([]*model.Question, error) {
	u, err := us.QuestionRepository.FindAll(u)
	if err != nil {
		return nil, err
	}

	return us.QuestionPresenter.ResponseQuestions(u), nil
}
