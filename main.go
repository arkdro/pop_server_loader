package main

import (
	"github.com/asdf/db_load/loader"
	"flag"
)

func main() {
	var file = flag.String("infile", "", "input file")
	var hst = flag.String("host", "localhost", "db host")
	var port = flag.Int("port", 3306, "db port")
	var user = flag.String("user", "root", "db user")
	var password = flag.String("password", "root", "db password")
	var database = flag.String("database", "pdata", "database")
	db_data := loader.Db_data{
		Host: *hst,
		Port: *port,
		User: *user,
		Password: *password,
		Database: *database,
	}
	loader.Load(*file, db_data)
}

