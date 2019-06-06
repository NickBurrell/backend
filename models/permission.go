package models

type Permission struct {
	PermissionID int `gorm:""`
	Name         string
	Description  string
}
