package repository

import (
	"github.com/jinzhu/gorm"
	"github.com/Shuto-san/4babe_api/domain/model"
)

type questionRepository struct {
	db *gorm.DB
}

type QuestionRepository interface {
	FindAll(u []*model.Question) ([]*model.Question, error)
}

func NewQuestionRepository(db *gorm.DB) QuestionRepository {
	return &questionRepository{db}
}

func (ur *questionRepository) FindAll(u []*model.Question) ([]*model.Question, error) {
	err := ur.db.Find(&u).Error

	if err != nil {
		return nil, err
	}

	return u, nil
}
