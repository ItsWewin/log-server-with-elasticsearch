package store

import (
	"testing"
	"time"
)

func TestMysqlDB_Connect(t *testing.T) {
	dsn := "klooktest:wewin123@tcp(127.0.0.1:3306)/projectKDB"
	connMaxLifetime := time.Duration(8) * time.Hour
	maxIdleConns := 100
	maxOpenConns := 20
	mysqlDB := NewMySQLDB(dsn, connMaxLifetime, maxIdleConns, maxOpenConns)

	_, err := mysqlDB.Connect()
	if err != nil {
		t.Fatalf("DB Connect: %s", err)
	}

	t.Logf("db connet succeed")
}
