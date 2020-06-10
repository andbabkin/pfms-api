package storage

import (
	"os"

	_ "github.com/go-sql-driver/mysql" // initializes a driver for database/sql api
	"github.com/jmoiron/sqlx"
)

var (
	rdbx *sqlx.DB
)

// OpenConns opens a pool of connections to relational database
func OpenConns() error {
	var err error
	rdbx, err = sqlx.Open("mysql", os.Getenv("MYSQL_URL"))
	if err != nil {
		return err
	}
	rdbx.SetMaxOpenConns(140)
	rdbx.SetMaxIdleConns(5)

	return nil
}

// GetRDBx returns an instance of sqlx (service which extends database/sql api)
func GetRDBx() *sqlx.DB {
	return rdbx
}
