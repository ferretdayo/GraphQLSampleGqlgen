package db

type Database struct {
	MainDB *DatabaseTarget
}

func (database *Database) Close() {
	database.MainDB.Close()
}
