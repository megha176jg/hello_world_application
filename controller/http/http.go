package http

import (
	"helloworld/service"
	"net/http"

	"bitbucket.org/junglee_games/getsetgo/monitoring"
	"github.com/gin-gonic/gin"
)

type HttpConfig struct {
	Port                string `yaml:"port"`
	TimeOutInSeconds    int    `yaml:"timeOutInSeconds"`
	MaxIdleConns        int    `yaml:"maxIdleConns"`
	MaxIdleConnsPerHost int    `yaml:"maxIdleConnsPerHost"`
}

type HttpController struct {
	config          HttpConfig
	monitoringAgent monitoring.Agent
	service         service.Service
}

func NewHttpController(config HttpConfig, ma monitoring.Agent, s service.Service) *HttpController {
	return &HttpController{
		config:          config,
		monitoringAgent: ma,
		service:         s,
	}
}

func (c *HttpController) StartListening() error {

	router := gin.Default()
	router.GET("/greet", func(ctx *gin.Context) {
		c.service.Greet(ctx)
	})
	err := http.ListenAndServe(":"+c.config.Port, router)
	if err != nil {
		return err
	}
	return nil
}
