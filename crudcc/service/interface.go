package service

import "github.com/shaurya-zopsmart/crudcc/models"

type User interface {
	InsertUsr(user models.User) (models.User, error)
	GetAllUsrs() ([]*models.User, error)
	GetUsrById(Id int) (models.User, error)
	UpdateUsr(id int, user models.User) error
	DeleteUsrById(Id int) error
	// GetEmail(Email string) (bool, error)
}
