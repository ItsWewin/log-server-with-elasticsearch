package main

import (
	"database/sql"
	"esServer/esParser/store"
	"log"
	"time"
)

var ProjectKDB *sql.DB

func main() {
	err := initDB()
	if err != nil {
		log.Fatalf("init db error: %s", err)
	}

	//var wg sync.WaitGroup
}

func initDB() error {
	var (
		dsn = "klooktest:wewin123@tcp(127.0.0.1:3306)/projectKDB"
		connMaxLifetime = time.Duration(8) * time.Hour
		maxIdleConns = 100
		maxOpenConns = 20
		err error
	)

	mysqlDB := store.NewMySQLDB(dsn, connMaxLifetime, maxIdleConns, maxOpenConns)
	ProjectKDB, err = mysqlDB.Connect()
	if err != nil {
		log.Fatalf("init db error: %s", err)
	}

	return nil
}
