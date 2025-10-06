package util_datasource

import (
	"github.com/jmoiron/sqlx"
)

type datasourceConnection struct {
	config datasourceBaseConfig

	provider datasourceProvider

	maxAttempt int

	db *sqlx.DB
}

/*
	Set config
*/
/*
	you can use BaseConfig function
	or
	datasourceBaseConfig struct
*/
func (d *datasourceConnection) SetConfig(config datasourceBaseConfig) *datasourceConnection {

	d.config = config
	return d
}

/*
	Set provider
*/
/*
	you can use Use[Provider Name]Provider function
*/
func (d *datasourceConnection) SetProvider(provider datasourceProvider) *datasourceConnection {

	d.provider = provider
	return d
}

/*
	Set max attempt
*/
/*
	max attempt for retry connect after error
*/
func (d *datasourceConnection) SetMaxAttempt(value int) *datasourceConnection {

	d.maxAttempt = value
	return d
}

/*
	Open datasource connection
*/
/*
	open connection by provider DSN and max attempt
*/
func (d *datasourceConnection) OpenConnection() error {

	db, err := d.provider.OpenConnection(d.provider.GetDSN(d.config), d.maxAttempt)
	if err != nil {
		return err
	}

	d.db = db

	return nil
}

/*
	Get connection
*/
/*
	if connection is nil will auto call OpenConnection function
*/
func (d *datasourceConnection) GetConn() *sqlx.DB {

	if d.db == nil {
		d.OpenConnection()
	}

	return d.db
}

/*
	Clean up
*/
/*
	if connection is nil simply return nil,
	if connection exist return error from Close function
*/
func (d *datasourceConnection) Cleanup() error {

	if d.db != nil {
		return nil
	}
	return d.db.Close()
}

/*
	Ping
*/
/*
	return error from Ping function
*/
func (d *datasourceConnection) Ping() error {

	return d.db.Ping()
}

/*
	Just ping
*/
/*
	return bool from is ping equal nil,
	if nil return true,
	if not nil return false
*/
func (d *datasourceConnection) JustPing() bool {

	return d.db.Ping() == nil
}

/*
	Sanitize query
*/
/*
	sanitize query by provider placeholder
	useful if you write query in mysql format and want to change to specific provider
*/
/*
	Caution
*/
/*
	!!! if your query is write in specific provider format,
	no need to use this function
*/
func (d *datasourceConnection) SanitizeQuery(rawQuery string) string {

	return queryReplacePlaceholder(rawQuery, d.provider.GetPlaceholder())
}
