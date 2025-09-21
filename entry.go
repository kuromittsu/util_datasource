package util_datasource

func CreateNewConnection() *DatasourceConnection {
	return &DatasourceConnection{}
}

func BaseConfig(host, port, username, password string) DatasourceBaseConfig {
	return DatasourceBaseConfig{
		Host:     host,
		Port:     port,
		Username: username,
		Password: password,
	}
}

func UseMysqlProvider(databaseName string) *MysqlProvider {
	return &MysqlProvider{
		DatabaseName: databaseName,
	}
}

func UsePostgresProvider(serviceName string) *PostgresProvider {
	return &PostgresProvider{
		ServiceName: serviceName,
	}
}

func UseOracleProvider(serviceName string) *OracleProvider {
	return &OracleProvider{
		ServiceName: serviceName,
	}
}
