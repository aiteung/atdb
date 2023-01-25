package atdb

import (
	"database/sql"
	"fmt"

	_ "github.com/microsoft/go-mssqldb"
)

func MssqlConnect(mssqlconn DBInfo) (db *sql.DB) {
	db, err := sql.Open("sqlserver", mssqlconn.DBString)
	if err != nil {
		fmt.Printf("MssqlConnect: %v\n", err)
	}
	return db

}
