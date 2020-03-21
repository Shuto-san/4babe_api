package presenter

import "github.com/Shuto-san/4babe_api/domain/model"

type questionPresenter struct {
}

type QuestionPresenter interface {
	ResponseQuestions(us []*model.Question) []*model.Question
}

func NewQuestionPresenter() QuestionPresenter {
	return &questionPresenter{}
}

func (up *questionPresenter) ResponseQuestions(us []*model.Question) []*model.Question {
	return us
}
