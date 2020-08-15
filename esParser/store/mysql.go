package store

import (
	"database/sql"
	"errors"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

type MysqlDB struct {
	DSN string `json:"dsn"`
	ConnMaxLifetime time.Duration `json:"conn_max_lifetime"`
	MaxIdleConns int `json:"max_idle_conns"`
	MaxOpenConns int `json:"max_open_conns"`
}

func NewMySQLDB(dsn string, connMaxLifetime time.Duration,  maxIdleConns, maxOpenConns int) (*MysqlDB) {
	m := &MysqlDB{}
	m.DSN = dsn
	m.ConnMaxLifetime = connMaxLifetime
	m.MaxIdleConns = maxIdleConns
	m.MaxOpenConns = maxOpenConns

	return m
}

func (m *MysqlDB) Connect() (*sql.DB, error) {
	pool, err := sql.Open("mysql", m.DSN)
	if err != nil {
		return nil, errors.New(err.Error())
	}

	pool.SetConnMaxLifetime(m.ConnMaxLifetime)
	pool.SetMaxIdleConns(m.MaxIdleConns)
	pool.SetMaxIdleConns(m.MaxOpenConns)

	if err := pool.Ping(); err != nil {
		return nil, errors.New(err.Error())
	}

	return pool, nil
}
