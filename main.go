package util_datasource

func CreateNewConnection() *DatasourceConnection {
	return &DatasourceConnection{}
}

func BaseConfig(host, port, username, password string) datasourceBaseConfig {
	return datasourceBaseConfig{
		Host:     host,
		Port:     port,
		Username: username,
		Password: password,
	}
}

func UseMysqlProvider(databaseName string) *mysqlProvider {
	return &mysqlProvider{
		DatabaseName: databaseName,
	}
}

func UsePostgresProvider(databaseName string) *postgresProvider {
	return &postgresProvider{
		DatabaseName: databaseName,
	}
}

func UseOracleProvider(serviceName string) *oracleProvider {
	return &oracleProvider{
		ServiceName: serviceName,
	}
}
