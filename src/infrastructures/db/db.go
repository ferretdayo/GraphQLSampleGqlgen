package db

type DataBase struct {
	MainDB *DBTarget
}

func (database *DataBase) Close() {
	database.MainDB.Close()
}
