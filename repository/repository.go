package repository

type RepositoryCreator interface {
	NewApplicationRepository() ApplicationRepository
	NewUserRepository() UserRepository
}

// TransactionalRepositoryCreator creates repositories which could hold the
// transaction by caller. Caller could commit/rollback after receiving some
// specified event, e.g. error happened.
type TransactionalRepositoryCreator interface {
	NewApplicationTxHolderRepository() ApplicationTxHolderRepository
	NewUserTxHolderRepository() UserTxHolderRepository
}

type Repository interface {
	RepositoryCreator
	TransactionalRepositoryCreator
}
