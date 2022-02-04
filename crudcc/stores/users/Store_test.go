package users

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"reflect"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/shaurya-zopsmart/user-curd/crudcc/models"
)

func Newmock() (*sql.DB, sqlmock.Sqlmock) {
	db, mock, err := sqlmock.New()

	if err != nil {
		log.Fatalf("Error %s in opening the datase ", err)
	}
	return db, mock
}

func NewMock1() (*sql.DB, sqlmock.Sqlmock) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		log.Printf("Error %s in opening the databse connection", err)
	}
	return db, mock
}

func Test_GetUserById(t *testing.T) {
	db, mock := Newmock()
	N := New(db)
	testcase := []struct {
		desc   string
		id     int
		experr error
		expout *models.User
		mock   []interface{}
	}{
		{
			desc:   "case:1",
			id:     5,
			experr: nil,
			expout: &models.User{Id: 5, Name: "Shaurya", Email: "berbreik@gmail.com", Phone: "9131346359", Age: 34},
			mock: []interface{}{
				mock.ExpectQuery("SELECT Id,Name,Email,Phone,Age FROM user WHERE Id=?").WithArgs(2).WillReturnRows(sqlmock.NewRows([]string{"Id", "Name", "Email", "Phone", "Age"}).AddRow(5, "Shaurya", "berbreik@gmail.com", "9131346359", 34)),
			},
		},
		{
			desc:   "case:2",
			id:     43,
			experr: nil,
			expout: &models.User{},
			mock: []interface{}{
				mock.ExpectQuery("SELECT Id,Name,Email,Phone,Age FROM user WHERE Id=?").WithArgs(43).WillReturnRows(sqlmock.NewRows([]string{"Id", "Name", "Email", "Phone", "Age"}).AddRow(0, "", "", "", 0)),
			},
		},
		{
			desc:   "case : 3",
			id:     8,
			experr: errors.New("error no row"),
			expout: &models.User{},
			mock: []interface{}{
				mock.ExpectQuery("SELECT Id,Name,Email,Phone,Age FROM user WHERE Id=?").WithArgs(8).WillReturnError(errors.New("error no row")),
			},
		},
	}
	for _, tcs := range testcase {
		resp, err := N.GetUserById(tcs.id)
		if !reflect.DeepEqual(resp, tcs.expout) {
			t.Errorf("expected %v got %v\n", tcs.expout, resp)
		}
		if err != nil && !errors.Is(err, tcs.experr) {
			t.Errorf("Error : expected %v got %v\n", tcs.experr, err)
		}
		if err != nil {
			fmt.Printf("expected %v got %v\n", tcs.experr, err)
		}
		fmt.Printf("expected %v got %v \n", tcs.experr, err)
	}
}

func Test_InsertUser(t *testing.T) {
	db, mock := NewMock1()
	N := New(db)
	testcase := []struct {
		desc   string
		value  models.User
		experr error
		expout models.User
		mock   []interface{}
	}{
		{
			desc:   "case 1",
			value:  models.User{Id: 3, Name: "Shaurya", Email: "berbreik@gmail.com", Phone: "9131346359", Age: 56},
			experr: nil,
			expout: models.User{},
			mock: []interface{}{
				mock.ExpectQuery("INSERT INTO user (Id,Name,Email,Phone,Age) Values(?,?,?,?,?)").WithArgs(3, "Shaurya", "berbreik@gmail.com", "9131346359", 56).WillReturnRows(sqlmock.NewRows([]string{"Id", "Name", "Email", "Phone", "Age"}).AddRow(3, "Shaurya", "berbreik@gmail.com", "9131346359", 56)),
			},
		},
		{
			desc:   "Case 2",
			value:  models.User{Id: 1, Name: "shaurya1", Email: "berbreik@gmail.com", Phone: "9131346359", Age: 54},
			experr: errors.New("fail to execute the query"),
			expout: models.User{},
			mock: []interface{}{
				mock.ExpectExec("INSERT INTO user (Id,Name,Email,Phone,Age) Values(?,?,?,?,?)").WithArgs(1, "shaurya1", "berbreik1@gmail.com", "9131346359", 54).WillReturnError(errors.New("fail to execute the query")),
			},
		},
	}
	for _, tcs := range testcase {
		resp, err := N.CreateUser(tcs.value)
		if !reflect.DeepEqual(resp, tcs.expout) {
			t.Errorf("Expected %v and got %v\n", tcs.expout, resp)
		}
		if err != nil && errors.Is(err, tcs.experr) {
			t.Errorf("Expected %v and got %v\n", tcs.experr, err)
		}
		if err != nil {
			fmt.Printf("Expected %v and got %v \n", tcs.experr, err)
		}
	}

}

func Test_GetAllUser(t *testing.T) {
	db, mock := Newmock()
	N := New(db)

	testcase := []struct {
		desc   string
		experr error
		expout []models.User
		mock   []interface{}
	}{
		{
			desc:   "case 1",
			experr: nil,
			expout: []models.User{
				{Id: 5, Name: "Shisui", Email: "shisui@gmail.com", Phone: "9131346359", Age: 23},
				{Id: 6, Name: "Shisui1", Email: "shisui1@gmail.com", Phone: "91313463591", Age: 231},
			},
			mock: []interface{}{
				mock.ExpectQuery("SELECT * FROM user").WillReturnRows(sqlmock.NewRows([]string{"id", "name", "email", "phone", "age"}).AddRow(
					5, "Shisui", "shisui@gmail.com", "9131346359", 23,
				).AddRow(
					6, "Shisui1", "shisui1@gmail.com", "91313463591", 231,
				)),
			},
		},
		{
			desc:   "case 2",
			experr: errors.New("fail to fetch the data"),
			expout: []models.User{},
			mock: []interface{}{
				mock.ExpectQuery("SELECT * FROM user").WillReturnError(errors.New("fail to fetch the data")),
			},
		},
	}
	for _, tcs := range testcase {
		resp, err := N.GetAllUsers()
		if !reflect.DeepEqual(resp, tcs.expout) {
			t.Errorf("Expected %v and got %v\n", tcs.expout, resp)
		}
		if err != nil && !errors.Is(err, tcs.experr) {
			t.Errorf("expected error %v but got %v\n", tcs.experr, err)
		}
		fmt.Printf("expected error %v but got %v\n", tcs.experr, err)
	}
}

func Test_update(t *testing.T) {
	db, mock := Newmock()
	N := New(db)

	testcase := []struct {
		desc   string
		value  models.User
		id     int
		experr error
		expout models.User
		mock   []interface{}
	}{
		{
			desc:   "case 1",
			value:  models.User{Id: 4, Name: "Itachi uchiha", Email: "Itachiuchia@gmail.com", Phone: "9131346359", Age: 13},
			id:     4,
			experr: nil,
			mock: []interface{}{
				mock.ExpectQuery("UPDATE user SET name=? , email=?, phone = ? , age = ? WHERE id = ?").WithArgs("Itachi uchiha", "Itachiuchiha@gmail.com", "9131346359", 13, 4).WillReturnRows(sqlmock.NewRows([]string{"id", "name", "email", "phone", "age"}).AddRow(4, "Itachi uchiha", "Itachiuchiha@gmail.com", "9131346359", 13)),
			},
		},
		{
			desc:   "case 2",
			value:  models.User{Id: 4, Name: "Itachi uchiha", Email: "Itachiuchia@gmail.com", Phone: "9131346359", Age: 13},
			id:     4,
			experr: errors.New("fail to execute the query"),
			mock: []interface{}{
				mock.ExpectQuery("UPDATE user SET name=? , email=?, phone = ? , age = ? WHERE id = ?").WithArgs("Itachi uchiha", "Itachiuchiha@gmail.com", "9131346359", 13, 9).WillReturnError(errors.New("fail to execute the query")),
			},
		},
	}
	for _, tcs := range testcase {
		resp, err := N.UpdateUser(tcs.id, tcs.value)
		if !reflect.DeepEqual(resp, tcs.value) {
			t.Errorf("Expected %v but got %v\n", tcs.value, resp)
		}
		if err != nil && errors.Is(err, tcs.experr) {
			t.Errorf("expected error %v but got %v\n", tcs.experr, err)
		}

	}
}

func Test_deletebyid(t *testing.T) {
	db, mock := Newmock()
	N := New(db)
	testcase := []struct {
		desc   string
		experr error
		id     int
		mock   []interface{}
	}{
		{
			desc:   "case 1",
			experr: errors.New("wrong query"),
			id:     6,
			mock: []interface{}{
				mock.ExpectQuery("DELETE FROM user WHERE id = ?").WithArgs(6).WillReturnError(errors.New("wrong query")),
			},
		},
	}
	for _, tcs := range testcase {
		err := N.DeleteUserById(tcs.id)
		if err != nil && err != tcs.experr {
			t.Errorf("expected %v but got %v\n", tcs.experr, err)
		}
	}
}
