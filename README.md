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
datasourceConnection

// open connection by provider DSN and max attempt
func(*datasourceConnection) OpenConnection() error

// getting connection db, if connection is nil will auto call OpenConnection function
func(*datasourceConnection) GetConn() *sqlx.DB

// close current connection
func(*datasourceConnection) Cleanup() error

// testing connection
func(*datasourceConnection) Ping() error

// testing connection but return bool
func(*datasourceConnection) JustPing() bool

// filter / sanitize query base on provider
func(*datasourceConnection) SanitizeQuery(rawQuery string) string

// set config for connection
func(*datasourceConnection) SetConfig(config datasourceBaseConfig) *datasourceConnection

// set provider for connection
func(*datasourceConnection) SetProvider(provider datasourceProvider) *datasourceConnection

// set max attempt for auto reconnect
func(*datasourceConnection) SetMaxAttempt(value int) *datasourceConnection
```

### Functions

- Creating new connection

```go
CreateNewConnection() *datasourceConnection
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
import (
  "fmt"
	"os"

  uds "github.com/kuromittsu/util_datasource"
)

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
