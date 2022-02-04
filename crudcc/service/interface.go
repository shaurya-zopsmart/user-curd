package service

import "github.com/shaurya-zopsmart/user-curd/crudcc/models"

type User interface {
	Create(user models.User) (models.User, error)
	GetAllUsrs() ([]*models.User, error)
	GetUsrById(Id int) (models.User, error)
	UpdateUsr(id int, user models.User) error
	DeleteUsrById(Id int) error
}
