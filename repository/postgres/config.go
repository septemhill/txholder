package postgres

import (
	"gorm.io/gorm"
)

type postgresConfig struct {
	db *gorm.DB
	tx *gorm.DB
}

func (config *postgresConfig) Db(db *gorm.DB) {
	config.db = db
}

func (config *postgresConfig) Tx(tx *gorm.DB) {
	config.tx = tx
}
