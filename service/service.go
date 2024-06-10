package service

import (
	"helloworld/repository" // Add the import statement for the "repository" package

	ac "helloworld/business/config"

	"bitbucket.org/junglee_games/getsetgo/monitoring"
	"github.com/gin-gonic/gin"
)

type Service interface {
	Greet(ctx *gin.Context)
}

type service struct {
	monitoringAgent monitoring.Agent
	repo            repository.Repository
	appConf         ac.AppConf
}

func NewService(repo repository.Repository, ma monitoring.Agent, a ac.AppConf) *service {
	return &service{
		repo:            repo,
		monitoringAgent: ma,
		appConf:         a,
	}
}

func (s *service) Greet(ctx *gin.Context) {

	name, ok := ctx.GetQuery("name")
	if ok != true {
		ctx.JSON(400, gin.H{"error": "name is required"})
		return
	}
	title, err := s.repo.GetTitle(name)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, gin.H{"message": "Hello, " + name + " " + title,
		"age": "Min age is:" + s.appConf.MinAge})
}
