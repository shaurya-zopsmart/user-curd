package stores

import "github.com/shaurya-zopsmart/user-curd/crudcc/models"

type Store interface {
	Create(usr models.User) (models.User, error)
	GetAllUsers() ([]*models.User, error)
	GetUserById(Id int) (*models.User, error)
	UpdateUser(Id int, usr models.User) (models.User, error)
	DeleteUserById(Id int) error
}
