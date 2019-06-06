package models

type Role struct {
	PermissionID int `gorm:""`
	Name         string
	Description  string
}
