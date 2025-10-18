package util_datasource

import (
	"fmt"
	"strings"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/sijms/go-ora/v2"
)

type oracleProvider struct {
	ServiceName string
}

func (d *oracleProvider) GetPlaceholder() string {
	return ":"
}

func (d *oracleProvider) GetDSN(config datasourceBaseConfig) string {
	return fmt.Sprintf("oracle://%s:%s@%s:%s/%s",
		config.Username,
		config.Password,
		config.Host,
		config.Port,
		d.ServiceName,
	)
}

func (d *oracleProvider) OpenConnection(dsn string, maxAttempts int) (*sqlx.DB, error) {

	var errorList []string

	for i := 0; i < maxAttempts; i++ {
		conn, err := sqlx.Connect("oracle", dsn)
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
