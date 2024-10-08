package {{cookiecutter.domain_name}}

import (
	"context"

	"github.com/google/uuid"
)

type Repository interface {
	Create(ctx context.Context, domain {{cookiecutter.domain_struct}}) (*uuid.UUID, error)
}
