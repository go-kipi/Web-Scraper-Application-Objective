package db

import (
	"fmt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type DB struct {
	Db *gorm.DB
}

func NewDb() DbI {
	return &DB{}
}

type DbI interface {
	InitDb() error
	SaveToSqlite(data interface{}) error
	GetFromSqlite() (interface{}, error)
	CountFromSqlite() (int64, error)
}

func (d *DB) InitDb() error {
	if db, err := gorm.Open(sqlite.Open("localScrape.db"), &gorm.Config{}); err != nil {
		return err
	} else {
		if client, err := db.DB(); err != nil {
			return err
		} else {
			query := `CREATE TABLE IF NOT EXISTS pokemons(
			   id integer PRIMARY KEY AUTOINCREMENT NOT NULL,
			   name           TEXT    NOT NULL,
			   image_url            TEXT     NOT NULL
				);`
			if _, err := client.Exec(query); err != nil {
				return err
			}
			d.Db = db
			return nil
		}
	}
}

func (d *DB) CountFromSqlite() (int64, error) {
	var count int64
	if err := d.Db.Table("pokemons").Count(&count); err.Error != nil {
		return 0, fmt.Errorf("can not Count pokemons to sqlite: %v", err)
	}
	return count, nil
}

func (d *DB) SaveToSqlite(data interface{}) error {
	if count, err := d.CountFromSqlite(); err != nil {
		return err
	} else if count == 0 {
		if err := d.Db.Table("pokemons").Save(data).Error; err != nil {
			return fmt.Errorf("can not save pokemons to sqlite: %v", err)
		}
	} else {
		return fmt.Errorf("database already exists")
	}

	return nil
}

func (d *DB) GetFromSqlite() (interface{}, error) {
	var data []struct {
		Name     string
		ImageUrl string `gorm:"column:image_url"`
	}

	if err := d.Db.Table("pokemons").Select("name").Scan(&data).Error; err != nil {
		return nil, fmt.Errorf("can not get pokemons to sqlite: %v", err)
	}
	return data, nil
}
