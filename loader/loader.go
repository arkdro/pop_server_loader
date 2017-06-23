package loader

func Load(file string, db_data Db_data) {
	fd, err := xlsx.OpenFile(file)
	if err != nil {
	}
	db, db_err := open_database(db_data)
	if db_err != nil {
	}
	process_data(fd, db)
}

func open_database(db_data Db_data) (string, error) {
	return "", nil
}

func process_data(fd *xlsx.File, db string) {
}

