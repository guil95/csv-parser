package cli

import "context"

type CLI interface {
	Run(ctx context.Context) error
}
