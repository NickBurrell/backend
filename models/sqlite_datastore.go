package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type SQLiteDatastore struct {
	db *gorm.DB
}

func NewSQLiteDatastore(path string) (*SQLiteDatastore, error) {
	db, err := gorm.Open("sqlite3", path)
	return &SQLiteDatastore{db}, err
}

func (d *SQLiteDatastore) GetUser(user *User) error {
	return d.db.First(user).Error
}

func (d *SQLiteDatastore) AddUser(user *User) error {
	return d.db.Create(user).Error
}

func (d *SQLiteDatastore) UpdateUser(user *User) error {
	return d.db.Update(user).Error
}

func (d *SQLiteDatastore) DeleteUser(user *User) error {
	return d.db.Delete(user).Error
}
