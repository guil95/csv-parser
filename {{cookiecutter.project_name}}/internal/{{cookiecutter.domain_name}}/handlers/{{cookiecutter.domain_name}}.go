package handlers

import (
	"github.com/{{cookiecutter.github_account}}/{{cookiecutter.github_repo_name}}/internal/{{cookiecutter.domain_name}}/usecases"
	"github.com/gin-gonic/gin"
)

type DomainHandler interface {
	Create(c *gin.Context)
}

type domainHandler struct {
	uc usecases.{{cookiecutter.domain_use_case}}
}

func NewDomainHandler(uc usecases.{{cookiecutter.domain_use_case}}) DomainHandler {
	return &domainHandler{
		uc: uc,
	}
}
