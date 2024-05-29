package postgresql

import (
	"database/sql"
	"fmt"
	"site/internal/init"
)

var DB *sql.DB

func Connect() {
	connStr := fmt.Sprintf("%s %s %s %s %s %s", init.DBconn.Host, init.DBconn.Port, init.DBconn.User, init.DBconn.Password, init.DBconn.DbName, init.DBconn.SSLMode)
	var err error
	DB, err = sql.Open("postgres", connStr)
	if err != nil {
		fmt.Printf("Error conection to DB. db.go:13: Error is %v", err)
	}
}
