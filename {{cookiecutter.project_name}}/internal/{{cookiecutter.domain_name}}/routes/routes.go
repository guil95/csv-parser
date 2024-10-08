package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/{{cookiecutter.github_account}}/{{cookiecutter.github_repo_name}}/internal/{{cookiecutter.domain_name}}/handlers"
)

type Opts struct {
	DomainHandler handlers.DomainHandler
}

func SetupRoutes(router *gin.Engine, opts Opts) {
	domainGroup := router.Group("/{{cookiecutter.domain_name}}s")
	{
		domainGroup.POST("", opts.DomainHandler.Create)
	}
}
