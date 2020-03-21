package router

import (
	"github.com/Shuto-san/4babe_api/interface/controller"
	"github.com/gin-gonic/gin"
)

func NewRouter(g *gin.Engine, c controller.AppController) *gin.Context {

	g.GET("/questions", func(context *gin.Context) error { return c.GetQuestions(context) })

	return g
}
