package usecases

import (
	"context"
	"github.com/{{cookiecutter.github_account}}/{{cookiecutter.github_repo_name}}/internal/{{cookiecutter.domain_name}}"
	"github.com/google/uuid"
	"log/slog"
)

func (uc {{cookiecutter.domain_use_case_pvt}}) Create(ctx context.Context, {{cookiecutter.domain_name}} {{cookiecutter.domain_name}}.{{cookiecutter.domain_struct}}) (*uuid.UUID, error) {
	ticketID, err := uc.repo.Create(ctx, {{cookiecutter.domain_name}})
	if err != nil {
		slog.Error("create {{cookiecutter.domain_name}} error", "error", err)
		return nil, err
	}

	return ticketID, nil
}
