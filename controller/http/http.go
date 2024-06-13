package http

import (
	"helloworld/service"

	"github.com/kataras/iris/v12"
)

type HttpConfig struct {
	Port                string
	TimeOutInSeconds    int
	MaxIdleConns        int
	MaxIdleConnsPerHost int
}

type HttpController struct {
	config  HttpConfig
	service service.Service
}

func NewHttpController(config HttpConfig, s service.Service) *HttpController {
	return &HttpController{
		config:  config,
		service: s,
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
