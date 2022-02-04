package users

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/shaurya-zopsmart/user-curd/crudcc/models"
)

type DataStore struct {
	db *sql.DB
}

func New(db *sql.DB) DataStore {
	return DataStore{db}
}
func (u *DataStore) CreateUser(usr models.User) (models.User, error) {

	Query := "INSERT INTO user (Id,Name,Email,Phone,Age) Values(?,?,?,?,?)"
	res, err := u.db.Exec(Query, usr.Id, usr.Name, usr.Email, usr.Phone, usr.Age)

	if err != nil {
		return usr, errors.New("fail to execute the query")
	}
	lid_64, err := res.LastInsertId()
	if err != nil {
		return usr, errors.New("fail to fetch the last inserted id")
	}
	lid := int(lid_64)
	resp, err := u.GetUserById(lid)

	if err != nil {
		return usr, errors.New("fail to fetch the last inserted id")
	}
	return resp, nil
}

func (u *DataStore) GetAllUsers() ([]*models.User, error) {
	var users []*models.User
	row, err := u.db.Query("SELECT * FROM user")
	if err != nil {
		return []*models.User{}, errors.New("fail to fetch the data")
	}
	defer row.Close()
	for row.Next() {
		var usr models.User
		_ = row.Scan(&usr.Id, &usr.Name, &usr.Email, &usr.Phone, &usr.Age)

		users = append(users, &usr)
	}

	return users, nil

}

func (u *DataStore) GetUserById(Id int) (models.User, error) {
	var solo models.User

	row := u.db.QueryRow("SELECT Id,Name,Email,Phone,Age FROM user WHERE Id=?", Id)
	err := row.Scan(&solo.Id, &solo.Name, &solo.Email, &solo.Phone, &solo.Age)
	if err != nil {
		return solo, errors.New("error no row")
	}
	return solo, nil
}

func (u *DataStore) UpdateUser(id int, value models.User) (models.User, error) {

	query := "Update User Set "
	var arg []interface{}

	if value.Name != "" {

		query = query + "Name = ?,"
		arg = append(arg, value.Name)
	}

	if value.Email != "" {

		query = query + "Email = ?,"
		arg = append(arg, value.Email)

	}

	if value.Phone != "" {

		query = query + "Phone = ?,"
		arg = append(arg, value.Phone)

	}

	if value.Age != 0 {

		query = query + "Age = ?,"
		arg = append(arg, value.Age)

	}
	query = query[:len(query)-1]
	query = query + " where Id = ?"
	arg = append(arg, id)

	fmt.Println(query, arg)
	_, err := u.db.Exec(query, arg...)

	if err != nil {
		return models.User{
			Id:    id,
			Name:  "",
			Phone: "",
			Email: "",
			Age:   0,
		}, errors.New("error in the given query")
	}

	user := models.User{}

	user, err = u.GetUserById(id)

	if err != nil {

		return models.User{
			Id:    id,
			Name:  "",
			Phone: "",
			Email: "",
			Age:   0,
		}, errors.New("error in the given query")
	}

	return user, nil

}

func (u *DataStore) DeleteUserById(id int) error {

	_, err := u.db.Exec("DELETE FROM user WHERE id = ? ", id)
	if err != nil {
		return errors.New("wrong query")
	}
	return nil
}
