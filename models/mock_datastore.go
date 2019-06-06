package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type MockDatastore struct {
	db *gorm.DB
}

func NewMockDatastore(path string) (*MockDatastore, error) {
	db, err := gorm.Open(":memory:", path)
	return &MockDatastore{db}, err
}
