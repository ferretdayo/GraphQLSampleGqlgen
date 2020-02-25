package db

import "github.com/jinzhu/gorm"

type DatabaseTarget struct {
	Master      *gorm.DB
	ReadReplica *gorm.DB
}

func (databaseTarget *DatabaseTarget) Close() {
	databaseTarget.Master.Close()
	databaseTarget.ReadReplica.Close()
}
