package route

import (
	"stockbit/app/controller"
	"stockbit/app/repository"
	"stockbit/app/service"

	"github.com/gin-gonic/gin"
)

func Init(router *gin.Engine) {
	ctrl := initController()
	group := router.Group("/api/v1/")

	filmGroup := group.Group("film")
	filmGroup.Use()
	{
		filmGroup.GET("", ctrl.GetFilmsByParams)
		filmGroup.GET("/:id", ctrl.GetFilmsByID)
	}
}

func initController() *controller.Controller {
	return controller.Init(*service.Init(repository.Init()))
}
