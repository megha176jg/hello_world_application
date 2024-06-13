package service

import (
	"helloworld/repository" // Add the import statement for the "repository" package

	ac "helloworld/business/config"

	"bitbucket.org/junglee_games/getsetgo/configs"
	"bitbucket.org/junglee_games/getsetgo/logger"
	"bitbucket.org/junglee_games/getsetgo/monitoring/monitoringfactory"
	"github.com/kataras/iris/v12"
)

type Service interface {
	Greet(ctx iris.Context)
}

type service struct {
	repo             repository.Repository
	monitoringConfig *configs.DefaultMonitoringConfig
	appConf          ac.AppConf
}

type result struct {
	Name string
	Age  string
}

func NewService(repo repository.Repository, m *configs.DefaultMonitoringConfig, a ac.AppConf) *service {
	return &service{
		repo:             repo,
		monitoringConfig: m,
		appConf:          a,
	}
}

func (s *service) Greet(ctx iris.Context) {

	agent, err := monitoringfactory.GetMonitoringAgent(s.monitoringConfig)
	if err != nil {
		panic(err)
	}
	defer agent.StartTransaction("Greeting").End()
	firstname := ctx.URLParam("name")
	title, err := s.repo.GetTitle(firstname)
	if err != nil {
		logger.Error(ctx, "unable to get the title")
		return
	}
	ctx.StatusCode(200)
	ctx.JSON(result{
		Name: firstname + " " + title,
		Age:  s.appConf.MinAge,
	})
}
