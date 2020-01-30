package db

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
)

type Mysql struct {
	User            string
	Password        string
	Host            string
	Name            string
	MasterPort      string
	ReadReplicaPort string
	Option          string
}

func NewMysql() *Mysql {
	return &Mysql{
		User:            os.Getenv("DB_USER"),
		Password:        os.Getenv("DB_PASSWORD"),
		Host:            os.Getenv("DB_HOST"),
		Name:            os.Getenv("DB_Name"),
		MasterPort:      os.Getenv("DB_MASTER_PORT"),
		ReadReplicaPort: os.Getenv("DB_READ_REPLICA_PORT"),
		Option:          os.Getenv("DB_OPTION"),
	}
}

func (mysql *Mysql) Open() *DataBase {
	masterDB, err := gorm.Open("mysql", fmt.Sprintf("%s:%s@(%s:%s)/%s?%s", mysql.User, mysql.Password, mysql.Host, mysql.MasterPort, mysql.Name, mysql.Option))
	if err != nil {
		panic(err)
	}

	readReplicaDB, err := gorm.Open("mysql", fmt.Sprintf("%s:%s@(%s:%s)/%s?%s", mysql.User, mysql.Password, mysql.Host, mysql.ReadReplicaPort, mysql.Name, mysql.Option))
	if err != nil {
		panic(err)
	}

	return &DataBase{
		MainDB: &DBTarget{
			Master:      masterDB,
			ReadReplica: readReplicaDB,
		},
	}
}
