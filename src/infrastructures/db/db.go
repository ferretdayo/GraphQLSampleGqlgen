package db

type Database struct {
	MainDB *DBTarget
}

func (database *Database) Close() {
	database.MainDB.Close()
}
