package user

import (
	"fmt"
	"net/mail"

	"github.com/shaurya-zopsmart/crudcc/models"
	"github.com/shaurya-zopsmart/crudcc/stores"
)

type UsrDtl struct {
	u stores.Store
}

func New(usrstr stores.Store) UsrDtl {
	return UsrDtl{u: usrstr}
}

func (usr UsrDtl) GetEmail(email string) (bool, error) {
	_, err := mail.ParseAddress(email)
	if err != nil {
		return false, err
	}
	usrs, err := usr.u.GetAllUsers()
	for _, ux := range usrs {
		if ux.Email == email {
			return false, err
		}
	}
	return true, nil
}

func (usr UsrDtl) GetAllUsrs() ([]models.User, error) {
	var us []models.User
	res, err := usr.u.GetAllUsers()
	if err != nil {
		return us, fmt.Errorf("%v", err)
	}
	us = res
	return us, nil

}

func (usr UsrDtl) GetUsrById(id int) (models.User, error) {
	var us models.User
	res, err := usr.u.GetUserById(id)
	if err != nil {
		return us, fmt.Errorf("%v", err)
	}
	us.Id = res.Id
	us.Name = res.Name
	us.Email = res.Email
	us.Phone = res.Phone
	us.Age = res.Age
	return us, nil

}

func (usr UsrDtl) InsertUsr(name, email, phone string, age int) (models.User, error) {
	user := models.User{0, name, email, phone, age}
	var nwsr models.User
	var err error
	user.Id, _ = usr.u.InsertUser(user)
	if err != nil {
		return nwsr, err
	}
	nwsr.Id = user.Id
	nwsr.Name = user.Name
	nwsr.Phone = user.Phone
	nwsr.Email = user.Email
	nwsr.Age = user.Age
	return nwsr, err

}
func (usr UsrDtl) DeleteUsrById(id int) error {
	err := usr.u.DeleteUserById(id)
	if err != nil {
		return fmt.Errorf("%v", err)
	}
	return nil
}

func (usr UsrDtl) UpdateUsr(usrdata models.User, u models.User) (models.User, error) {
	err := usr.u.UpdateUser(u)
	u.Id = usrdata.Id
	if u.Name != "" {
		usrdata.Name = u.Name
	}
	valid, _ := usr.GetEmail(u.Email)
	if valid {
		usrdata.Email = u.Email
	}
	if u.Phone != "" {
		usrdata.Phone = u.Phone
	}
	if u.Age != 0 {
		usrdata.Age = u.Age
	}
	return usrdata, err

}
