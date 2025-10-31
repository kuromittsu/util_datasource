package util_datasource

import (
	"github.com/jmoiron/sqlx"
)

type datasourceProvider interface {

	/*
		Get DSN with base config
	*/
	GetDSN(config datasourceBaseConfig) string

	/*
		Open provider connection
	*/
	OpenConnection(dsn string, maxAttempts int) (*sqlx.DB, error)

	/*
		Get placeholder
	*/
	GetPlaceholder() string

	/*
		Get provider name
	*/
	GetProviderName() string
}
