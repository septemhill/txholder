package postgres

import (
	"txholder/repository"

	"gorm.io/gorm"
)

type tx gorm.DB

func (t tx) Apply(conf repository.Config) {
	conf.Tx((*gorm.DB)(&t))
}

func WithTx(t *gorm.DB) repository.Option {
	return tx(*t)
}

type db gorm.DB

func (d db) Apply(conf repository.Config) {
	conf.Db((*gorm.DB)(&d))
}

func WithDb(d *gorm.DB) repository.Option {
	return db(*d)
}
