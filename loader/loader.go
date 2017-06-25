package loader

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/tealeg/xlsx"
	"fmt"
	"log"
)

func Load(file string, db_data Db_data) {
	fd, err := xlsx.OpenFile(file)
	if err != nil {
		log.Fatal("can not open file: ", err)
	}
	db := open_database(db_data)
	process_data(fd, db)
}

func open_database(db_data Db_data) *sql.DB {
	db, err := sql.Open("mysql", "user:password@/dbname")
	if err != nil {
		log.Fatal("can not connect to the database: ", err)
	}
	return db
}

func process_data(fd *xlsx.File, db *sql.DB) {
	for _, sheet := range fd.Sheets {
		for _, row := range sheet.Rows {
			for _, cell := range row.Cells {
				text, _ := cell.String()
				fmt.Printf("%s\n", text)
			}
		}
	}
}
