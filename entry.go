package util_datasource

func CreateNewConnection() *datasourceConnection {
	return &datasourceConnection{}
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

func UsePostgresProvider(serviceName string) *postgresProvider {
	return &postgresProvider{
		ServiceName: serviceName,
	}
}

func UseOracleProvider(serviceName string) *oracleProvider {
	return &oracleProvider{
		ServiceName: serviceName,
	}
}
