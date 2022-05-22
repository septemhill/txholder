package postgres

import (
	repository "github.com/septemhill/txholder/repository"

	"github.com/jmoiron/sqlx"
)

type tx sqlx.Tx

func (t tx) Apply(conf repository.Config) {
	conf.Tx((*sqlx.Tx)(&t))
}

func WithTx(t *sqlx.Tx) repository.Option {
	return tx(*t)
}

type db sqlx.DB

func (d db) Apply(conf repository.Config) {
	conf.Db((*sqlx.DB)(&d))
}

func WithDb(d *sqlx.DB) repository.Option {
	return db(*d)
}
