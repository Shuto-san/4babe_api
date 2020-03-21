package presenter

import "github.com/Shuto-san/4babe_api/domain/model"

type QuestionPresenter interface {
	ResponseQuestions(u []*model.Question) []*model.Question
}
