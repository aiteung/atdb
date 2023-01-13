package atdb

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func MariaConnect(mariaconn DBInfo) (db *sql.DB) {
	db, err := sql.Open("mysql", mariaconn.DBString+mariaconn.DBName)
	if err != nil {
		fmt.Printf("MariaConnect: %v\n", err)
	}
	return db

}
