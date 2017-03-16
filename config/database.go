package config

import (
	"flag"
	"log"

	"upper.io/db.v3/lib/sqlbuilder"
	"upper.io/db.v3/sqlite"
)

var DBFile *string
var DB *sqlbuilder.Database

func DBParams() {
	envDBFile := envString("DB_FILE", "")
	DBFile = flag.String("db.file", ":"+envDBFile, "SQLite db file")

	setting := sqlite.ConnectionURL{
		Database: envDBFile,
	}
	//var err error
	db, err := sqlite.Open(setting)
	if err != nil {
		log.Panic(err)
	}
	if db != nil {
		DB = &db
	}
}