package models

type User struct {
	Email       string
	Username    string
	Password    string
	IsVerified  bool
	Roles       []Role
	Permissions []Permission
}
