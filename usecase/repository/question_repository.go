package repository

import "github.com/Shuto-san/4babe_api/domain/model"

type QuestionRepository interface {
	FindAll(u []*model.Question) ([]*model.Question, error)
}
