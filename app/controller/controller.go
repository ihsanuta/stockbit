package controller

import (
	"stockbit/app/service"
)

type Controller struct {
	service service.Service
}

func Init(service service.Service) *Controller {
	return &Controller{
		service: service,
	}
}
