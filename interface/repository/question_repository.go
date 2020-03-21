package repository

import (
	"github.com/Shuto-san/4babe_api/domain/model"
	"github.com/gomodule/redigo/redis"
	"github.com/jinzhu/gorm"
)

type questionRepository struct {
	db   *gorm.DB
	conn redis.Conn
}

type QuestionRepository interface {
	FindAll(u []*model.Question) ([]*model.Question, error)
}

func NewQuestionRepository(db *gorm.DB, conn redis.Conn) QuestionRepository {
	return &questionRepository{db, conn}
}

func (ur *questionRepository) FindAll(u []*model.Question) ([]*model.Question, error) {
	err := ur.db.Find(&u).Error

	if err != nil {
		return nil, err
	}

	ur.conn.Do("SET", "go-redis", "test")

	return u, nil
}
