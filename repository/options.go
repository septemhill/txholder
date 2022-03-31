package repository

import "gorm.io/gorm"

//go:generate mockery --name=Config --inpackage --case underscore
type Config interface {
	Db(*gorm.DB)
	Tx(*gorm.DB)
}

//go:generate mockery --name=Option --inpackage --case underscore
type Option interface {
	Apply(Config)
}
