package service

import (
	"helloworld/repository" // Add the import statement for the "repository" package

	ac "helloworld/business/config"

	"bitbucket.org/junglee_games/getsetgo/instrumenting/newrelic"
	"bitbucket.org/junglee_games/getsetgo/logger"
	"bitbucket.org/junglee_games/getsetgo/monitoring"
	"github.com/kataras/iris/v12"
	nrf "github.com/newrelic/go-agent/v3/newrelic"
)

type Service interface {
	Greet(ctx iris.Context)
}

type service struct {
	monitoringAgent monitoring.Agent
	repo            repository.Repository
	appConf         ac.AppConf
}

type result struct {
	Name string
	Age  string
}

func NewService(repo repository.Repository, ma monitoring.Agent, a ac.AppConf) *service {
	return &service{
		repo:            repo,
		monitoringAgent: ma,
		appConf:         a,
	}
}

func (s *service) Greet(ctx iris.Context) {
	txnObject, err := newrelic.GetNewrelicTxn(ctx)
	if err != nil {
		logger.Error(ctx, "unable to load get newrelic txnobject")
	}
	newCtx := nrf.NewContext(ctx.Request().Context(), txnObject)
	defer nrf.FromContext(newCtx).StartSegment("Greeting").End()
	firstname := ctx.URLParam("name")
	title, _ := s.repo.GetTitle(firstname)
	// if err != nil {
	// }
	ctx.StatusCode(200)
	ctx.JSON(result{
		Name: firstname + title,
		Age:  s.appConf.MinAge,
	})

}
