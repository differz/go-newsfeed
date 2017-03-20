package config

import (
	"log"
	"upper.io/db.v3/lib/sqlbuilder"
	"upper.io/db.v3/sqlite"
)

var DB sqlbuilder.Database

func DBParams() {
	envDBFile := envString("DB_FILE", "")

	setting := sqlite.ConnectionURL{
		Database: envDBFile,
	}
	//var err error
	db, err := sqlite.Open(setting)
	if err != nil {
		log.Panic(err)
	}
	if db != nil {
		DB = db
	}
}
