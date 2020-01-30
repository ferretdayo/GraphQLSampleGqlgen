package db

import "github.com/jinzhu/gorm"

type DBTarget struct {
	Master      *gorm.DB
	ReadReplica *gorm.DB
}

func (dbTarget *DBTarget) Close() {
	dbTarget.Master.Close()
	dbTarget.ReadReplica.Close()
}
