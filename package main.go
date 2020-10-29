package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	fmt.Println("Go Mysqldatabase")

	
	db, err := sql.Open("mysql", "username:Mahesh143@amma@tc(127.0.0.1:3306)/demo")

	
	if err != nil {
		panic(err.Error())
	}


	defer db.Close()

}