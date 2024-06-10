package controller

import (
	"helloworld/controller/http"
	"helloworld/service"

	"bitbucket.org/junglee_games/getsetgo/monitoring"
)

type Controller interface {
	StartListening() error
}

type Config struct {
	HTTP http.HttpConfig
}

func NewController(c Config, ma monitoring.Agent, s service.Service) Controller {
	return http.NewHttpController(c.HTTP, ma, s)
}
