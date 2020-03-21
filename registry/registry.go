package registry

import (
	"github.com/Shuto-san/4babe_api/interface/controller"
	"github.com/gomodule/redigo/redis"
	"github.com/jinzhu/gorm"
)

type registry struct {
	db   *gorm.DB
	conn redis.Conn
}

type Registry interface {
	NewAppController() controller.AppController
}

func NewRegistry(db *gorm.DB, conn redis.Conn) Registry {
	return &registry{db, conn}
}

func (r *registry) NewAppController() controller.AppController {
	return r.NewQuestionController()
}
