package main

import (
	"GoProjects/price/conf"
	"GoProjects/price/error"
	"database/sql"
	"fmt"
)

// Db подключение к базе
func Db() *sql.DB {
	conf := conf.Db
	connStr := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", conf.User, conf.Pass, conf.Host, conf.Port, conf.Name)

	db, err := sql.Open("mysql", connStr)
	error.CheckNil(err)

	return db
}
