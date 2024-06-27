package service

import (
	"helloworld/repository" // Add the import statement for the "repository" package

	ac "helloworld/business/config"

	"bitbucket.org/junglee_games/getsetgo/configs"
	"github.com/gin-gonic/gin"
)

type Service interface {
	Greet(ctx *gin.Context)
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

// @Summary		Greet name
// @Description	get title by name
// @Accept		*/*
// @Produce		json
// @Param		name	query		string	true	"Some Name"
// @Success		200		{string}	string	"ok"
// @Router		/greet/ [get]
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
