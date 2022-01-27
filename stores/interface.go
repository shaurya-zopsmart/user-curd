package stores

import "github.com/shaurya-zopsmart/crudcc/models"

type Store interface {
	InsertUser(models.User) (int, error)
	GetAllUsers() ([]models.User, error)
	GetUserById(Id int) (models.User, error)
	UpdateUser(models.User) error
	DeleteUserById(Id int) error
}
