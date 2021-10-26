package database

import (
	"database/sql"
	"shop/config"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func LoadDatabase() error {
	var err error
	dbconfig := config.Config.DBUsername + ":" + config.Config.DBPassword + "@" + config.Config.DBProtocal + "(" + config.Config.DBAddress + ")/" + config.Config.DBName

	DB, err = sql.Open(config.Config.DBType, dbconfig)
	if err != nil {
		return err
	}

	if err = DB.Ping(); err != nil || DB == nil {
		return err
	}

	return nil
}
