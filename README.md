# Util Datasource

## Install

```
go get github.com/kuromittsu/util_datasource
```

## Available sources

- [Mysql](#mysql)
- [Postgres](#postgres)
- [Oracle](#oracle)

## Usage

### Struct

- Datasource Connection

```go
DatasourceConnection

// open connection by provider DSN and max attempt
func(*DatasourceConnection) OpenConnection() error

// getting connection db, if connection is nil will auto call OpenConnection function
func(*DatasourceConnection) GetConn() *sqlx.DB

// close current connection
func(*DatasourceConnection) Cleanup() error

// testing connection
func(*DatasourceConnection) Ping() error

// testing connection but return bool
func(*DatasourceConnection) JustPing() bool

// filter / sanitize query base on provider
func(*DatasourceConnection) SanitizeQuery(rawQuery string) string

// set config for connection
func(*DatasourceConnection) SetConfig(config datasourceBaseConfig) *DatasourceConnection

// set provider for connection
func(*DatasourceConnection) SetProvider(provider datasourceProvider) *DatasourceConnection

// get provider name
func(*DatasourceConnection) GetProviderName() string

// set max attempt for auto reconnect
func(*DatasourceConnection) SetMaxAttempt(value int) *DatasourceConnection
```

### Functions

- Creating new connection

```go
CreateNewConnection() *DatasourceConnection
```

- Base Config

```go
BaseConfig(
  "host",
  "port",
  "username",
  "password"
) datasourceBaseConfig
```

other functions see [sources](#sources)

### Example

```go
// import uds "github.com/kuromittsu/util_datasource"

// create new connection
conn1 := uds.CreateNewConnection()

// set config (host, port, username, password)
conn1.SetConfig(uds.BaseConfig(
  "localhost",
  "5432",
  "your_username",
  "your_password",
))

// set provider
// for example, set provider to postgres provider and required parameter is database name
conn1.SetProvider(uds.UsePostgresProvider(
  "your_database_name"
))

// set max attempt (int)
// for example, set max attempt to 3
conn1.SetMaxAttempt(3)

// call function to open connection
if err := conn1.OpenConnection(); err != nil {
  fmt.Printf("error while opening connection | %v \n", err)
  os.Exit(0)
}

// just ping to check current connection status
fmt.Printf("ping: %v\n", conn1.JustPing())

// call cleanup after connection finish used
if err := conn1.Cleanup(); err != nil {
  fmt.Printf("error while cleanup | %v \n", err)
  os.Exit(0)
}
```

### Sources

#### Mysql

```go
UseMysqlProvider(
  "database_name"
)
```

#### Postgres

```go
UsePostgresProvider(
  "database_name"
)
```

#### Oracle

```go
UseOracleProvider(
  "service_name"
)
```
