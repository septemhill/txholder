package postgres

import "github.com/jmoiron/sqlx"

type postgresConfig struct {
	db *sqlx.DB
	tx *sqlx.Tx
}

func (config *postgresConfig) Db(db *sqlx.DB) {
	config.db = db
}

func (config *postgresConfig) Tx(tx *sqlx.Tx) {
	config.tx = tx
}
