package main

import (
	"database/sql"

	"log"

	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:@(172.17.0.4:3306)/")
	if err != nil {
		log.Fatalf("Error while connecting to database %v", err)
	} else {
		fmt.Print("Connected to DBServer successfuly !")
	}
	defer db.Close()
}
