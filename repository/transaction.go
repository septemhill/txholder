package repository

// TransactionHolderRepository
//go:generate mockery --name=TransactionHolderRepository --inpackage --case underscore
type TransactionHolderRepository interface {
	CommitableRepository
	RollbackableRepository
}

// CommitableRepository makes transaction holder repostiory could commit.
//go:generate mockery --name=CommitableRepository --inpackage --case underscore
type CommitableRepository interface {
	Commit() error
}

// RollbackableRepository makes transaction holder repostiory could rollback.
//go:generate mockery --name=RollbackableRepository --inpackage --case underscore
type RollbackableRepository interface {
	Rollback() error
}
