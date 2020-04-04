package db

import (
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
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
		Name:            os.Getenv("DB_NAME"),
		MasterPort:      os.Getenv("DB_MASTER_PORT"),
		ReadReplicaPort: os.Getenv("DB_READ_REPLICA_PORT"),
		Option:          os.Getenv("DB_OPTION"),
	}
}

func (mysql *Mysql) Open() *Database {
	masterDB, err := gorm.Open("mysql", fmt.Sprintf("%s:%s@(%s:%s)/%s?%s", mysql.User, mysql.Password, mysql.Host, mysql.MasterPort, mysql.Name, mysql.Option))
	if err != nil {
		panic(err)
	}

	readReplicaDB, err := gorm.Open("mysql", fmt.Sprintf("%s:%s@(%s:%s)/%s?%s", mysql.User, mysql.Password, mysql.Host, mysql.ReadReplicaPort, mysql.Name, mysql.Option))
	if err != nil {
		panic(err)
	}

	return &Database{
		MainDB: &DatabaseTarget{
			Master:      masterDB,
			ReadReplica: readReplicaDB,
		},
	}
}
