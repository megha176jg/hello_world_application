package http

import (
	"helloworld/service"

	"bitbucket.org/junglee_games/getsetgo/monitoring"
	"github.com/kataras/iris/v12"
)

type HttpConfig struct {
	Port                string
	TimeOutInSeconds    int
	MaxIdleConns        int
	MaxIdleConnsPerHost int
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
	router := iris.New()
	router.Get("/greet", func(ctx iris.Context) {
		c.service.Greet(ctx)
	})
	// err := http.ListenAndServe(":"+c.config.Port, router)
	router.Run(iris.Addr(":" + c.config.Port))
	// if err != nil {
	// 	return err
	// }
	return nil
}
