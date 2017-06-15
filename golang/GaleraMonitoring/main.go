package main

import (
	"database/sql"

	"log"

	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var n *sql.DB

func main() {

	n1, err := sql.Open("mysql", "root:@(172.17.0.2:3306)/")
	if err != nil {
		log.Fatalf("Error while connectiong to N1 %v", err)
	}
	defer n1.Close()

	n2, err := sql.Open("mysql", "root:@(172.17.0.3:3306)/")
	if err != nil {
		log.Fatalf("Error while connectiong to N3 %v", err)
	}
	defer n2.Close()

	n3, err := sql.Open("mysql", "root:@(172.17.0.4:3306)/")
	if err != nil {
		log.Fatalf("Error while connectiong to N3 %v", err)
	}
	defer n3.Close()

	getVersion(n1)
	getVersion(n2)
	getVersion(n3)

	getServerUUID(n1)
	getServerUUID(n2)
	getServerUUID(n3)

}

func getVersion(n *sql.DB) {
	rows, err := n.Query("select version()")
	if err != nil {
		log.Fatal(err)
	}
	for rows.Next() {
		var version string
		err = rows.Scan(&version)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(version)
	}
}

func getServerUUID(n *sql.DB) {
	rows, err := n.Query("SHOW VARIABLES LIKE \"%server_uuid%\"")
	if err != nil {
		log.Fatal(err)
	}
	for rows.Next() {

		var server string
		var uid string

		err = rows.Scan(&server, &uid)
		if err != nil {
			log.Fatal(err)
		}
		//fmt.Println(server)
		fmt.Println(uid)
	}
}
