package users

import (
	"database/sql"
	"fmt"

	"github.com/shaurya-zopsmart/crudcc/models"
)

type DataStore struct {
	db *sql.DB
}

func New(db *sql.DB) DataStore {
	return DataStore{db: db}
}
func (u *DataStore) InsertUser(usr models.User) (int, error) {

	Query := "INSERT INTO user(Id,Name,Email,Phone,Age) Values(?,?,?,?,?)"
	res, err := u.db.Exec(Query, usr.Id, usr.Name, usr.Email, usr.Phone, usr.Age)

	if err != nil {
		return -1, fmt.Errorf("%v", err)
	}
	lastid, err := res.LastInsertId()
	if err != nil {
		return -1, fmt.Errorf("%v", err)

	}
	return int(lastid), nil
}

func (u *DataStore) GetAllUsers() ([]models.User, error) {
	var users []models.User
	row, _ := u.db.Query("SELECT * FROM user")
	defer row.Close()
	for row.Next() {
		var usr models.User
		if err := row.Scan(&usr.Id, &usr.Name, &usr.Email, &usr.Phone, &usr.Age); err != nil {
			fmt.Errorf("%v", err)
		}
		users = append(users, usr)
	}
	if err := row.Err(); err != nil {
		fmt.Errorf("%v", err)
	}
	return users, nil

}

func (u *DataStore) GetUserById(id int) (models.User, error) {
	var solo models.User
	if id < 1 {
		return solo, fmt.Errorf("not a valid id")
	}
	res := u.db.QueryRow("SELECT * FROM user WHERE id=? ", id)
	res.Scan(&solo.Id, &solo.Name, &solo.Email, &solo.Phone, &solo.Age)
	return solo, nil
}

func (u *DataStore) UpdateUser(usr models.User) error {
	if usr.Id < 1 {
		return fmt.Errorf("not a valid id less than 1")
	}
	_, err := u.db.Exec("UPDATE user SET name=? , email=?, phone = ? , age = ? WHERE id = ?", usr.Name, usr.Email, usr.Phone, usr.Age, usr.Id)
	if err != nil {
		return fmt.Errorf("%s", err)
	}
	return nil
}

func (u *DataStore) DeleteUserById(id int) error {
	if id < 1 {
		fmt.Errorf("invalid id")
	}
	_, err := u.db.Exec("DELETE FROM user WHERE id = ? ", id)
	if err != nil {
		return fmt.Errorf("%s", err)
	}
	return nil
}
