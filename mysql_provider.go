package util_datasource

import (
	"fmt"
	"strings"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type MysqlProvider struct {
	DatabaseName string
}

func (d *MysqlProvider) GetPlaceholder() string {
	return ""
}

func (d *MysqlProvider) GetDSN(config DatasourceBaseConfig) string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		config.Username,
		config.Password,
		config.Host,
		config.Port,
		d.DatabaseName,
	)
}

func (d *MysqlProvider) OpenConnection(dsn string, maxAttempts int) (*sqlx.DB, error) {

	var errorList []string

	for i := 0; i < maxAttempts; i++ {
		conn, err := sqlx.Connect("mysql", dsn)
		if err == nil {
			return conn, nil
		}

		errorList = append(errorList, err.Error())

		fmt.Printf("creating connection error | %v \n", err)
		fmt.Printf("try to reconnecting after %v \n", time.Second*time.Duration(i+1))

		time.Sleep(time.Second * time.Duration(i+1))
	}

	return nil, fmt.Errorf("result error | %v", strings.Join(errorList, ", "))
}
