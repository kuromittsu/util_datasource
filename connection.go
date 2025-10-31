package util_datasource

import (
	"errors"

	"github.com/jmoiron/sqlx"
)

type DatasourceConnection struct {
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
func (d *DatasourceConnection) SetConfig(config datasourceBaseConfig) *DatasourceConnection {

	d.config = config
	return d
}

/*
	Set provider
*/
/*
	you can use Use[Provider Name]Provider function
*/
func (d *DatasourceConnection) SetProvider(provider datasourceProvider) *DatasourceConnection {

	d.provider = provider
	return d
}

/*
	Get provider name
*/
/*
	Get provider name from provider
*/
func (d *DatasourceConnection) GetProviderName() string {

	if d.provider == nil {
		return ""
	}
	return d.provider.GetProviderName()
}

/*
	Set max attempt
*/
/*
	max attempt for retry connect after error
*/
func (d *DatasourceConnection) SetMaxAttempt(value int) *DatasourceConnection {

	d.maxAttempt = value
	return d
}

/*
	Open datasource connection
*/
/*
	open connection by provider DSN and max attempt
*/
func (d *DatasourceConnection) OpenConnection() error {

	if d.provider == nil {
		return errors.New("no provider yet, provider is required")
	}

	if d.maxAttempt < 1 {
		d.maxAttempt = 1
	}

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
func (d *DatasourceConnection) GetConn() *sqlx.DB {

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
func (d *DatasourceConnection) Cleanup() error {

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
func (d *DatasourceConnection) Ping() error {

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
func (d *DatasourceConnection) JustPing() bool {

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
func (d *DatasourceConnection) SanitizeQuery(rawQuery string) string {

	return queryReplacePlaceholder(rawQuery, d.provider.GetPlaceholder())
}
