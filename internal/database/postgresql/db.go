package postgresql

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"site/internal/config"
)

var DB *gorm.DB

var DBConn = config.DBConnection{Host: "localhost", Port: "5432", User: "postgres", Password: "Admin@@246", DbName: "site", SSLMode: "disable"}

func Connect() {
	connStr := fmt.Sprintf("%s %s %s %s %s %s", DBConn.Host, DBConn.Port, DBConn.User, DBConn.Password, DBConn.DbName, DBConn.SSLMode)
	var err error
	DB, err = gorm.Open(postgres.Open(connStr), &gorm.Config{})
	if err != nil {
		fmt.Printf("Error conection to DB. db.go:13: Error is %v", err)
	}
}
