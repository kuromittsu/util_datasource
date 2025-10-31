package util_datasource

import (
	"fmt"
	"strings"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type postgresProvider struct {
	DatabaseName string
}

func (d *postgresProvider) GetProviderName() string {
	return "postgres"
}

func (d *postgresProvider) GetPlaceholder() string {
	return "$"
}

func (d *postgresProvider) GetDSN(config datasourceBaseConfig) string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		config.Host,
		config.Port,
		config.Username,
		config.Password,
		d.DatabaseName,
	)
}

func (d *postgresProvider) OpenConnection(dsn string, maxAttempts int) (*sqlx.DB, error) {

	var errorList []string

	for i := 0; i < maxAttempts; i++ {
		conn, err := sqlx.Connect("postgres", dsn)
		if err == nil {
			return conn, nil
		}

		errorList = append(errorList, err.Error())

		fmt.Printf("creating connection error | %v \n", err)
		fmt.Printf("try to reconnecting after %v \n", time.Second*time.Duration(i+1))

		time.Sleep(time.Second * time.Duration(i+1))
	}

	return nil, fmt.Errorf("error result(s):\n%v", strings.Join(errorList, ",\n"))
}
