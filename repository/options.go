package repository

import "github.com/jmoiron/sqlx"

type Config interface {
	Db(*sqlx.DB)
	Tx(*sqlx.Tx)
}

type Option interface {
	Apply(Config)
}
