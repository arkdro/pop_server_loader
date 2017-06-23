package main

import (
	"github.com/asdf/db_load/loader"
)

func main() {
	file := ""
	db_data := loader.Db_data{
	}
	loader.Load(file, db_data)
}

