package bootstrap

import (
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"github.com/{{cookiecutter.github_account}}/{{cookiecutter.github_repo_name}}/internal/{{cookiecutter.domain_name}}/handlers"
	"github.com/{{cookiecutter.github_account}}/{{cookiecutter.github_repo_name}}/internal/{{cookiecutter.domain_name}}/routes"
	"github.com/{{cookiecutter.github_account}}/{{cookiecutter.github_repo_name}}/internal/{{cookiecutter.domain_name}}/storage"
	"github.com/{{cookiecutter.github_account}}/{{cookiecutter.github_repo_name}}/internal/{{cookiecutter.domain_name}}/usecases"
)

func Domain(router *gin.Engine, db *sqlx.DB) {
	repository := storage.NewPostgresRepository(db)
	uc := usecases.New{{cookiecutter.domain_use_case}}(repository)
	opts := routes.Opts{
		DomainHandler: handlers.NewDomainHandler(uc),
	}

	routes.SetupRoutes(router, opts)
}
