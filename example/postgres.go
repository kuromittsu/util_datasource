package main

import (
	"fmt"
	"os"

	"github.com/kuromittsu/util_datasource"
)

func testingPostgres() {

	conn := util_datasource.CreateNewConnection().SetConfig(util_datasource.BaseConfig(
		"localhost",
		"5432",
		"username",
		"password",
	)).SetProvider(util_datasource.UsePostgresProvider("database_name")).SetMaxAttempt(5)

	fmt.Println("opening connection ...")
	if err := conn.OpenConnection(); err != nil {
		fmt.Printf("error while opening connection | %v \n", err)
		os.Exit(0)
	}
	fmt.Println("open connection success")

	fmt.Printf("ping: %v\n", conn.JustPing())

	// add query operation here

	fmt.Println("cleanup ...")
	if err := conn.Cleanup(); err != nil {
		fmt.Printf("error while cleanup | %v \n", err)
		os.Exit(0)
	}
	fmt.Println("cleanup success")

}
