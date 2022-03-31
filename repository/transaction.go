package repository

// TransactionHolderRepository
type TransactionHolderRepository interface {
	CommitableRepository
	RollbackableRepository
}

// CommitableRepository makes transaction holder repostiory could commit.
type CommitableRepository interface {
	Commit() error
}

// RollbackableRepository makes transaction holder repostiory could rollback.
type RollbackableRepository interface {
	Rollback() error
}
