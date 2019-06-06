package models

type User struct {
	Email       string
	Username    string
	Password    string
	Roles       []Role
	Permissions []Permission
}
