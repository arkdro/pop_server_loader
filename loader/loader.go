package loader

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/tealeg/xlsx"
	"fmt"
	"log"
//	"time"
	"strconv"
)

const (
	country_idx = 2
	country_code_idx = 4
	cell_begin_idx = 5
	cell_end_idx = 19
)

type Row struct {
	country string
	country_code string
	data []string
}

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
	for idx_s, sheet := range fd.Sheets {
		if idx_s != 0 {
			continue
		}
		data_started_flag := false
		var years []string
		for _, sheet_row := range sheet.Rows {
			if data_started_flag {
				row := extract_data_from_row(sheet_row)
				log.Printf("row: %v\n", row)
				if valid_values(row) {
					store_row(row, db, years)
				}
			} else {
				if is_index_row(sheet_row) {
					years = fill_years(sheet_row)
					data_started_flag = true
				}
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

func extract_data_from_row(row *xlsx.Row) Row {
	country := row.Cells[country_idx].Value
	country_code := row.Cells[country_code_idx].Value
	data := row.Cells[cell_begin_idx : cell_end_idx]
	cells := make([]string, len(data))
	for i := range data {
		cells[i] = data[i].Value
	}
	res := Row{
		country: country,
		country_code: country_code,
		data: cells,
	}
	return res
}

func valid_values(row Row) bool {
	if is_empty(row.country) {
		return false
	}
	if is_empty(row.country_code) {
		return false
	}
	for _, x := range row.data {
		if is_empty(x) {
			return false
		}
		if !is_float_str(x) {
			return false
		}
	}
	return true
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

func store_row(row Row, db *sql.DB, years []string) {
	cmd := `insert into country_median_age
        (country, country_code, year, age)
        values (?, ?, ?, ?)`
	for i, age := range row.data {
		year := years[i]
		res, err := db.Exec(cmd, row.country, row.country_code, year, age)
		log.Println("row, res: ", res, ", err: ", err)
	}
}

