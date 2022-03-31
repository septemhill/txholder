package repository

import "gorm.io/gorm"

type Config interface {
	Db(*gorm.DB)
	Tx(*gorm.DB)
}

type Option interface {
	Apply(Config)
}
