package controller

import (
	"helloworld/controller/http"
	"helloworld/service"
)

type Controller interface {
	StartListening() error
}

type Config struct {
	HTTP http.HttpConfig
}

func NewController(c Config, s service.Service) Controller {
	return http.NewHttpController(c.HTTP, s)
}
