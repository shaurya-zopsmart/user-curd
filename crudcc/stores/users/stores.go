package users

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/shaurya-zopsmart/crudcc/models"
)

type DataStore struct {
	db *sql.DB
}

func New(db *sql.DB) DataStore {
	return DataStore{db}
}
func (u *DataStore) InsertUser(usr models.User) (models.User, error) {

	Query := "INSERT INTO user (Id,Name,Email,Phone,Age) Values(?,?,?,?,?)"
	res, err := u.db.Exec(Query, usr.Id, usr.Name, usr.Email, usr.Phone, usr.Age)

	if err != nil {
		return usr, errors.New("fail to execute the query")
	}
	_, err = res.LastInsertId()
	if err != nil {
		return usr, errors.New("fail to fetch the last inserted id")
	}
	return usr, nil
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
	// if err := row.Err(); err != nil {
	// 	return users, errors.New("fail to append fucntion or scan")
	// }
	return users, nil

}

func (u *DataStore) GetUserById(Id int) (*models.User, error) {
	var solo models.User

	row := u.db.QueryRow("SELECT Id,Name,Email,Phone,Age FROM user WHERE Id=?", Id)
	err := row.Scan(&solo.Id, &solo.Name, &solo.Email, &solo.Phone, &solo.Age)
	if err != nil {
		return &solo, errors.New("error no row")
	}
	return &solo, nil
}

func (u *DataStore) UpdateUser(Id int, usr models.User) (models.User, error) {

	// if id < 1 {
	// 	return usr, fmt.Errorf("not a valid id less than 1")
	// }
	_, err := u.db.Exec("UPDATE user SET name=? , email=?, phone = ? , age = ? WHERE id = ?", usr.Name, usr.Email, usr.Phone, usr.Age, Id)
	if err != nil {
		return usr, fmt.Errorf("fail to execute the query%s", err)
	}

	return usr, nil

}

func (u *DataStore) DeleteUserById(id int) error {

	_, err := u.db.Exec("DELETE FROM user WHERE id = ? ", id)
	if err != nil {
		return errors.New("wrong query")
	}
	return nil
}
