package types

import "database/sql"

var db *sql.DB

type ConnectDB interface {
	connect() string
	execute(string) (string, error)
	close() (string, error)
}

type Config struct {
	User     string
	Password string
	Net      string
	Addr     string
	Port     string
	DBName   string
}

func (c *Config) CreateConfig(driver string) {
	var dataConfig Config = *c
	db, err = sql.Open(driver, dataConfig)
}
