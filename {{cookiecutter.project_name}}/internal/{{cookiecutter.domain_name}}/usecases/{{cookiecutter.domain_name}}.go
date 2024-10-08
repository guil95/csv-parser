package usecases

import (
	"context"
	"github.com/{{cookiecutter.github_account}}/{{cookiecutter.github_repo_name}}/internal/{{cookiecutter.domain_name}}"
	"github.com/google/uuid"
)

type (
	{{cookiecutter.domain_use_case}} interface {
		Create(ctx context.Context, {{cookiecutter.domain_name}} {{cookiecutter.domain_name}}.{{cookiecutter.domain_struct}}) (uuid *uuid.UUID, err error)
	}

	{{cookiecutter.domain_use_case_pvt}} struct {
		repo {{cookiecutter.domain_name}}.Repository
	}
)

func New{{cookiecutter.domain_use_case}}(repo {{cookiecutter.domain_name}}.Repository) {{cookiecutter.domain_use_case}}  {
	return &{{cookiecutter.domain_use_case_pvt}}{repo: repo}
}
