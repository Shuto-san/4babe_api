package controller

import (
	"net/http"

	"github.com/Shuto-san/4babe_api/domain/model"
	"github.com/Shuto-san/4babe_api/usecase/interactor"
)

type questionController struct {
	questionInteractor interactor.QuestionInteractor
}

type QuestionController interface {
	GetQuestions(c Context) error
}

func NewQuestionController(us interactor.QuestionInteractor) QuestionController {
	return &questionController{us}
}

func (uc *questionController) GetQuestions(c Context) error {
	var u []*model.Question

	u, err := uc.questionInteractor.Get(u)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, u)
}
