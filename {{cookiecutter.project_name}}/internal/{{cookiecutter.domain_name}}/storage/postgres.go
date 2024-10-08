package storage

import (
	"context"
	"github.com/google/uuid"
	"github.com/{{cookiecutter.github_account}}/{{cookiecutter.github_repo_name}}/internal/contact"
	"github.com/jmoiron/sqlx"
)

type postgresRepository struct {
	db *sqlx.DB
}

func NewPostgresRepository(db *sqlx.DB) {{cookiecutter.domain_name}}.Repository {
	return &postgresRepository{db: db}
}

func (repo *postgresRepository) Create(ctx context.Context, domain {{cookiecutter.domain_name}}.{{cookiecutter.domain_struct}}) (*uuid.UUID, error) {
	return nil, nil
}
