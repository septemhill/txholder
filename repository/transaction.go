package repository

import "context"

type Transactional interface {
	Commit(ctx context.Context) error
	Rollback(ctx context.Context) error
}
