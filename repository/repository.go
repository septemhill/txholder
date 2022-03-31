package repository

//go:generate mockery --name=RepositoryCreator --inpackage --case underscore
type RepositoryCreator interface {
	NewApplicationRepository() ApplicationRepository
	NewUserRepository() UserRepository
}

// TransactionalRepositoryCreator creates repositories which could hold the
// transaction by caller. Caller could commit/rollback after receiving some
// specified event, e.g. error happened.
//go:generate mockery --name=TransactionalRepositoryCreator --inpackage --case underscore
type TransactionalRepositoryCreator interface {
	NewApplicationTxHolderRepository() ApplicationTxHolderRepository
	NewUserTxHolderRepository() UserTxHolderRepository
}

//go:generate mockery --name=Repository --inpackage --case underscore
type Repository interface {
	RepositoryCreator
	TransactionalRepositoryCreator
}
