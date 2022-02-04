package user

import (
	"errors"
	"reflect"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/shaurya-zopsmart/user-curd/crudcc/models"
	"github.com/shaurya-zopsmart/user-curd/crudcc/stores"
)

func Test_GetUserById(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockstore := stores.NewMockStore(ctrl)
	testuserservice := New(mockstore)

	testcase := []struct {
		desc     string
		id       int
		experr   error
		exp      models.User
		mockcall *gomock.Call
	}{
		{
			desc:   " case 1",
			id:     1,
			experr: nil,
			exp: models.User{
				Id:    1,
				Name:  "Shaurya",
				Email: "berbreik@gmail.com",
				Phone: "9131346359",
				Age:   23,
			},
			mockcall: mockstore.EXPECT().GetUserById(1).Return(&models.User{Id: 1, Name: "Shaurya", Email: "berbreik@gmail.com", Phone: "9131346359", Age: 23}, nil),
		},
		{
			desc:   "case 2",
			id:     2,
			experr: errors.New("cant fetch id"),
			exp:    models.User{},

			mockcall: mockstore.EXPECT().GetUserById(2).Return(&models.User{}, errors.New("cant fetch id")),
		},
		{
			desc:   "case 3",
			id:     -1,
			experr: errors.New("invalid id"),
		},
	}
	for _, tcs := range testcase {
		t.Run(tcs.desc, func(t *testing.T) {
			user, err := testuserservice.GetUsrById(tcs.id)
			if err != nil && !reflect.DeepEqual(tcs.exp, user) {
				t.Errorf("Expected : %v but got %v", tcs.exp, user)
			}
		})
	}
}

func Test_GetAllUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockstore := stores.NewMockStore(ctrl)
	testservice := New(mockstore)
	data := []models.User{
		{
			Id:    1,
			Name:  "Shaurya",
			Email: "as@gmail.com",
			Phone: "9876543230",
			Age:   23,
		},
		{
			Id:    2,
			Name:  "nit",
			Email: "nit@gmail.com",
			Phone: "9988776655",
			Age:   18,
		},
	}
	testcase := []struct {
		desc     string
		exp      []models.User
		mockcall *gomock.Call
	}{
		{
			desc:     "case 1",
			exp:      data,
			mockcall: mockstore.EXPECT().GetAllUsers().Return(data, nil),
		},
		{
			desc:     "case 2",
			exp:      []models.User{},
			mockcall: mockstore.EXPECT().GetAllUsers().Return([]models.User{}, errors.New("fail to execute service")),
		},
	}
	for _, tcs := range testcase {
		t.Run(tcs.desc, func(t *testing.T) {
			user, err := testservice.GetAllUsrs()
			if err != nil && !reflect.DeepEqual(tcs.exp, user) {
				t.Errorf("Expected %v but got %v", tcs.exp, err)
			}
		})
	}
}

func Test_update(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockstore := stores.NewMockStore(ctrl)
	testservice := New(mockstore)
	testuser := models.User{Name: "namaiwa", Email: "pitaparka@gmail.com", Phone: "9131346359", Age: 43}
	testcase := []struct {
		desc     string
		id       int
		exp      error
		mockcall *gomock.Call
	}{
		{
			desc:     "case 1",
			id:       1,
			exp:      nil,
			mockcall: mockstore.EXPECT().UpdateUser(1, testuser).Return(nil),
		},
		{
			desc:     "case 2",
			id:       4,
			exp:      errors.New("cant update"),
			mockcall: mockstore.EXPECT().UpdateUser(4, testuser).Return(errors.New("cant update")),
		},
	}
	for _, tcs := range testcase {
		t.Run(tcs.desc, func(t *testing.T) {
			err := testservice.UpdateUsr(tcs.id, testuser)

			if err != nil && errors.Is(tcs.exp, err) {
				t.Errorf("expected : %v but got %v", tcs.exp, err)
			}
		})
	}
}

func Test_deletebyid(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockstore := stores.NewMockStore(ctrl)
	testservice := New(mockstore)

	testcase := []struct {
		desc     string
		id       int
		exp      error
		mockcall *gomock.Call
	}{
		{
			desc:     "case 1",
			id:       2,
			exp:      nil,
			mockcall: mockstore.EXPECT().DeleteUserById(2).Return(nil),
		},
		{
			desc:     "case 2",
			id:       -1,
			exp:      errors.New("invalid id"),
			mockcall: mockstore.EXPECT().DeleteUserById(-1).Return(errors.New("invalid id")),
		},
	}
	for _, tcs := range testcase {
		t.Run(tcs.desc, func(t *testing.T) {
			err := testservice.DeleteUsrById(tcs.id)

			if err != nil && errors.Is(tcs.exp, err) {
				t.Errorf("expected %v but got %v\n", tcs.exp, err)

			}
		})
	}

}

func Test_InsertUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockstore := stores.NewMockStore(ctrl)
	testservice := New(mockstore)

	testUser := models.User{Id: 1, Name: "vv", Email: "vv@gmail.com", Phone: "9988776655", Age: 23}
	testcase := []struct {
		desc     string
		expout   models.User
		experr   error
		mockcall *gomock.Call
	}{
		{
			desc:     "case 1",
			expout:   models.User{Id: 1, Name: "vv", Email: "vv@gmail.com", Phone: "9988776655", Age: 23},
			experr:   nil,
			mockcall: mockstore.EXPECT().CreateUser(testUser).Return(models.User{Id: 1, Name: "vv", Email: "vv@gmail.com", Phone: "9988776655", Age: 23}, nil),
		},
		{
			desc:     "case 2",
			expout:   models.User{},
			experr:   errors.New("fail to execute"),
			mockcall: mockstore.EXPECT().CreateUser(models.User{}).Return(models.User{}, errors.New("fail to execute")),
		},
	}
	for _, tcs := range testcase {
		t.Run(tcs.desc, func(t *testing.T) {
			res, err := testservice.CreateUsr(tcs.expout)
			if !reflect.DeepEqual(tcs.expout, res) {
				t.Errorf("Expected %v but got %v", tcs.expout, res)
			}
			if err != nil && errors.Is(tcs.experr, err) {
				t.Errorf("Expeceted %v but got %v", tcs.experr, err)
			}
		})
	}
}
