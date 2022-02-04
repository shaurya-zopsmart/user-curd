package user

import (
	"errors"
	"log"

	"github.com/shaurya-zopsmart/user-curd/crudcc/models"
	"github.com/shaurya-zopsmart/user-curd/crudcc/stores"
)

type UsrDtl struct {
	u stores.Store
}

func New(usrstr stores.Store) *UsrDtl {
	return &UsrDtl{
		u: usrstr,
	}
}

func (usr UsrDtl) GetAllUsrs() ([]*models.User, error) {

	res, err := usr.u.GetAllUsers()
	if err != nil {
		return res, errors.New("fail to execute service")
	}
	return res, nil

}

func (usr UsrDtl) GetUsrById(id int) (models.User, error) {

	if validateid(id) {
		log.Println("invalid id")
	}
	user, err := usr.u.GetUserById(id)
	if err != nil {
		log.Println("cant fetch id")
	}
	return *user, nil

}

func (usr UsrDtl) Create(user models.User) (models.User, error) {
	if !validateid(user.Id) {
		return user, errors.New("error invalid id")
	}

	if !validatemail(user.Email) {
		return user, errors.New("error invalid email")
	}
	if !validp(user.Phone) {
		return user, errors.New("invalid phone")
	}
	seruser, err := usr.u.Create(user)
	if err != nil {
		return seruser, errors.New("fail to execute")
	}
	return seruser, nil

}
func (usr UsrDtl) DeleteUsrById(id int) error {

	if !validateid(id) {
		return errors.New("invalid id")
	}
	return usr.u.DeleteUserById(id)

}

func (usr UsrDtl) UpdateUsr(id int, user models.User) error {

	if !validateid(id) {
		return errors.New("errrors invalid id")
	}
	if user.Email != "" && !validatemail(user.Email) {
		return errors.New("error invaldi email")
	}
	if user.Phone != "" && !validp(user.Phone) {
		return errors.New("error invalid phone")
	}
	_, err := usr.u.UpdateUser(id, user)
	if err != nil {
		return errors.New("cant update")
	}
	return nil

}
