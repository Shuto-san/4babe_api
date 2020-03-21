package registry

import (
	"github.com/jinzhu/gorm"
	"github.com/Shuto-san/4babe_api/interface/controller"
)

type registry struct {
	db *gorm.DB
}

type Registry interface {
	NewAppController() controller.AppController
}

func NewRegistry(db *gorm.DB) Registry {
	return &registry{db}
}

func (r *registry) NewAppController() controller.AppController {
	return r.NewQuestionController()
}
