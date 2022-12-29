package config

import "github.com/test-framwork/types/"

func Config() {
	dbConfig := types.Config{
		User:     "postgre",
		Port:     "6543",
		Password: "1234",
	}
	var Db types.Config
	Db.CreateConfig()
}
