package service

import "github.com/shaurya-zopsmart/crudcc/models"

type User interface {
	InsertUsr(name, email, phone string, age int) (models.User, error)
	GetAllUsrs() ([]models.User, error)
	GetUsrById(Id int) (models.User, error)
	UpdateUsr(models.User, models.User) (models.User, error)
	DeleteUsrById(Id int) error
	GetEmail(Email string) (bool, error)
}
