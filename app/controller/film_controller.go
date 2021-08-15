package controller

import (
	"net/http"
	"stockbit/app/entity"

	"github.com/gin-gonic/gin"
)

func (c *Controller) GetFilmsByParams(ctx *gin.Context) {
	var params entity.ParamsFilm

	// query := ctx.Request.URL.Query()
	err := ctx.ShouldBindQuery(&params)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, map[string]interface{}{
			"status_code": http.StatusBadRequest,
			"message":     "params not found",
		})
		return
	}

	data, err := c.service.Film.GetFilmByParams(params)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, map[string]interface{}{
			"status_code": http.StatusBadRequest,
			"message":     err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, map[string]interface{}{
		"status_code": http.StatusOK,
		"message":     "success",
		"data":        data,
	})
}

func (c *Controller) GetFilmsByID(ctx *gin.Context) {
	filmID := ctx.Param("id")
	data, err := c.service.Film.GetByFilmByID(filmID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, map[string]interface{}{
			"status_code": http.StatusBadRequest,
			"message":     err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, map[string]interface{}{
		"status_code": http.StatusOK,
		"message":     "success",
		"data":        data,
	})
}
