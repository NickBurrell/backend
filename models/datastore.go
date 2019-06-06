package models

type Datastore interface {
	GetUser(*User) error
	AddUser(*User) error
	UpdateUser(*User) error
	DeleteUser(*User) error
}
