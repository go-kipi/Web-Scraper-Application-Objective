package db

import (
	"database/sql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type DB struct {
	Db *gorm.DB
}

func NewDb() DbI {
	return DB{}
}

type DbI interface {
	InitDb() (*sql.DB, error)
}

func (d DB) InitDb() (*sql.DB, error) {
	if db, err := gorm.Open(sqlite.Open("localScrape.db"), &gorm.Config{}); err != nil {
		return nil, err
	} else {
		if client, err := db.DB(); err != nil {
			return nil, err
		} else {
			return client, nil
		}
	}

}

//func (d DB) Get() error {
//	return d.Client.Ping()
//}
