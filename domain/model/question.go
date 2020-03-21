package model

import "time"

type Question struct {
	ID          uint      `gorm:"primary_key" json:"id"`
	QustionType string    `json:"question_type"`
	Question    string    `json:"question"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func (Question) TableName() string { return "questions" }
