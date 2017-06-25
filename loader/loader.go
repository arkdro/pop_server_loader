package loader

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/tealeg/xlsx"
	"fmt"
	"log"
)

func Load(file string, db_data Db_data) {
	fd := open_xls_file(file)
	db := open_database(db_data)
	process_data(fd, db)
}

func open_xls_file(file string) *xlsx.File {
	fd, err := xlsx.OpenFile(file)
	if err != nil {
		log.Fatal("can not open file: ", err)
	}
	return fd
}

func open_database(db_data Db_data) *sql.DB {
	dsn := build_dsn(db_data)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal("can not connect to the database: ", err)
	}
	return db
}

func build_dsn(db_data Db_data) string {
	dsn := db_data.User	+
		":" +
		db_data.Password +
		"@tcp(" +
		db_data.Host +
		":" +
		fmt.Sprint(db_data.Port) +
		")/" +
		db_data.Database
	return dsn
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
func is_index_row(row *xlsx.Row) bool {
	text := row.Cells[0].Value
	if text == "Index" {
		return true
	} else {
		return false
	}
}

func fill_years(row *xlsx.Row) []string {
	data := row.Cells[cell_begin_idx : cell_end_idx]
	cells := make([]string, len(data))
	for i := range data {
		cells[i] = data[i].Value
	}
	return cells
}

func is_empty(s string) bool {
	if s == "" {
		return true
	} else {
		return false
	}
}

func is_float_str(s string) bool {
	_, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return false
	} else {
		return true
	}
}

