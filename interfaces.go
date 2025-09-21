package util_datasource

import (
	"github.com/jmoiron/sqlx"
)

type DatasourceProvider interface {

	/*
		Get DSN with base config
	*/
	GetDSN(config DatasourceBaseConfig) string

	/*
		Open provider connection
	*/
	OpenConnection(dsn string, maxAttempts int) (*sqlx.DB, error)

	/*
		Get placeholder
	*/
	GetPlaceholder() string
}
