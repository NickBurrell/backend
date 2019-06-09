package models

type Role struct {
	RoleID      int `gorm:""`
	Name        string
	Description string
}
