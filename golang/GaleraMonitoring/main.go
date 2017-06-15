package main

import (
	"database/sql"
	"log"

	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {

	cnx := map[string]string{
		"n1": "root:@(172.17.0.2:3306)/",
		"n2": "root:@(172.17.0.3:3306)/",
		"n3": "root:@(172.17.0.4:3306)/",
	}

	dbList := map[string]*sql.DB{}

	for key, con := range cnx {
		db, err := sql.Open("mysql", con)
		if err != nil {
			log.Fatalf("Error while connecting to server %s : %v", key, err)
		}
		defer db.Close()
		dbList[key] = db
	}

	for srvName, db := range dbList {
		version, err := getVersion(db)
		if err != nil {
			log.Fatalf("Impossible to get version %v", err)
		}
		log.Printf("Serveur %s - version %s", srvName, version)
	}

	muid := map[string]string{}

	for srvName, db := range dbList {
		_, uid, err := getClusterStateUUID(db)
		if err != nil {
			log.Fatalf("Impossible to get uid %v", err)
		}
		muid[srvName] = uid
		log.Printf("%s %s", srvName, uid)
	}

	err := checkUID(muid)
	if err != nil {
		log.Fatalf("%s : %v", err, muid)
	}

}

func getVersion(db *sql.DB) (version string, err error) {

	q := "select version()"
	err = db.QueryRow(q).Scan(&version)

	return
}

func getClusterStateUUID(db *sql.DB) (srv, uid string, err error) {

	q := "SHOW GLOBAL STATUS LIKE 'wsrep_cluster_state_uuid'"
	err = db.QueryRow(q).Scan(&srv, &uid)

	return

}

func checkUID(uids map[string]string) error {

	lastUID := ""

	for srv, uid := range uids {
		if lastUID == "" {
			lastUID = uid
			continue
		}
		if lastUID == uid {
			continue
		}
		return fmt.Errorf("uid : %s of %s does not match", uid, srv)
	}

	return nil

}
