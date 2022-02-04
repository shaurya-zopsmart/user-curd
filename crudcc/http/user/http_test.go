package user

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
	"github.com/shaurya-zopsmart/user-curd/crudcc/models"
	"github.com/shaurya-zopsmart/user-curd/crudcc/service"
)

func Test_GetUserById(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockservice := service.NewMockUser(ctrl)
	mock := New(mockservice)

	testcase := []struct {
		desc    string
		id      string
		mock    []*gomock.Call
		experr  error
		expbody []byte
	}{
		{
			desc: "case 1",
			id:   "1",
			mock: []*gomock.Call{
				mockservice.EXPECT().GetUsrById(1).Return(models.User{
					Id:    1,
					Name:  "tea",
					Email: "tea@gmail.com",
					Phone: "9876543210",
					Age:   23,
				}, nil),
			},
			expbody: []byte(`{"Id":1,"Name":"tea","Email":"tea@gmail.com","Phone":"9876543210","Age":23}`),
		},
		{
			desc:    "Case 2: Failure case",
			id:      "1a",
			experr:  nil,
			expbody: []byte("Invalid id atoi"),
		},
		{
			desc:   "Case 3: Failure case - 2",
			id:     "100000000",
			experr: nil,
			mock: []*gomock.Call{
				mockservice.EXPECT().GetUsrById(100000000).Return(models.User{
					Id:    0,
					Name:  "",
					Email: "",
					Phone: "",
					Age:   0,
				}, errors.New("user id not found")),
			},
			expbody: []byte("user id not found"),
		},
	}
	for _, test := range testcase {
		req := httptest.NewRequest("GET", "/user/"+test.id, nil)
		res := httptest.NewRecorder()

		req = mux.SetURLVars(req, map[string]string{
			"id": test.id,
		})

		mock.UserwithID(res, req)

	}
}

func TestGetAllUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockservice := service.NewMockUser(ctrl)
	mock := New(mockservice)

	testcase := []struct {
		desc    string
		id      string
		mock    []*gomock.Call
		experr  error
		expbody []byte
	}{
		{
			desc:   "Case 1: Success case ",
			experr: nil,
			mock: []*gomock.Call{
				mockservice.EXPECT().GetAllUsrs().Return([]models.User{{
					Id:    1,
					Name:  "tea",
					Email: "tea@gmail.com",
					Phone: "9876543210",
					Age:   23,
				},
				}, nil),
			},
			expbody: []byte(`[{"Id":1,"Name":"tea","Email":"tea@gmail.com","Phone":"9876543210","Age":23}]`),
		},
		{
			desc:   "Failure case-1 ",
			experr: errors.New("Could not fetch users"),
			mock: []*gomock.Call{
				mockservice.EXPECT().GetAllUsrs().Return([]models.User{{
					Id:    2,
					Name:  "tejas",
					Email: "tejas@gmail.com",
					Phone: "9876543210",
					Age:   23,
				},
				}, errors.New("Could not fetch users")),
			},
			expbody: []byte("Could not fetch users"),
		},
	}
	for _, v := range testcase {
		t.Run(v.desc, func(t *testing.T) {
			r := httptest.NewRequest("GET", "/api/users/", nil)
			rw := httptest.NewRecorder()
			mock.GetAllUsers(rw, r)
			if rw.Body.String() != string(v.expbody) {
				t.Errorf("Expected %v Obtained %v", string(v.expbody), rw.Body.String())
			}
		})
	}

}

func TestUpdate(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockservice := service.NewMockUser(ctrl)
	mock := New(mockservice)

	testcase := []struct {
		desc    string
		user    models.User
		mock    []*gomock.Call
		experr  error
		expbody []byte
	}{
		{
			desc: "Case 1: Success case",
			user: models.User{
				Id:    1,
				Name:  "tea",
				Email: "tea@gmail.com",
				Phone: "9876543210",
				Age:   23,
			},
			mock: []*gomock.Call{
				mockservice.EXPECT().UpdateUsr(0, models.User{
					Id:    1,
					Name:  "tea",
					Email: "tea@gmail.com",
					Phone: "9876543210",
					Age:   23,
				}).Return(nil),
			},
			experr: nil,
		},
		{
			desc: "Case 1: Failure case -1",
			user: models.User{
				Id:    -2,
				Name:  "tejas",
				Email: "tejas@gmail.com",
				Phone: "9876543210",
				Age:   23,
			},
			experr:  errors.New("Invalid Id"),
			expbody: []byte("Invalid Id"),
		},
	}
	for _, v := range testcase {
		t.Run(v.desc, func(t *testing.T) {
			b, _ := json.Marshal(v.user)
			r := httptest.NewRequest("PUT", "/user"+string(rune(v.user.Id)), bytes.NewBuffer(b))
			rw := httptest.NewRecorder()
			mock.UpdateUser(rw, r)
			fmt.Println(rw.Body.String())
			if rw.Body.String() != string(v.expbody) {
				t.Errorf("Expected %v Obtained %v", string(v.expbody), rw.Body.String())
			}
		})
	}
}

func TestDelete(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockservice := service.NewMockUser(ctrl)
	mock := New(mockservice)

	testCases := []struct {
		desc     string
		id       string
		expecErr error
		mock     []*gomock.Call
		expecRes []byte
	}{
		{
			desc:     "Case 1: Success case",
			id:       "1",
			expecErr: nil,
			mock: []*gomock.Call{
				mockservice.EXPECT().DeleteUsrById(1).Return(nil),
			},
			expecRes: []byte("user deleted successfully"),
		},
		{
			desc:     "Case 2: Failure case - 1",
			id:       "-2",
			expecErr: errors.New("Invalid Id"),
			expecRes: []byte("Invalid Id"),
		},
		{
			desc:     "Case 3: Failure case - 2",
			id:       "0",
			expecErr: errors.New("error while deleting user"),
			mock: []*gomock.Call{
				mockservice.EXPECT().DeleteUsrById(0).Return(errors.New("error while deleting user")),
			},
			expecRes: []byte("error while deleting user"),
		},
	}
	for _, v := range testCases {
		t.Run(v.desc, func(t *testing.T) {
			r := httptest.NewRequest("DELETE", "/api/users/"+v.id, nil)
			rw := httptest.NewRecorder()
			mock.DeleteUser(rw, r)
			fmt.Println(rw.Body.String())
			if rw.Body.String() != string(v.expecRes) {
				t.Errorf("Expected %v Obtained %v", string(v.expecRes), rw.Body.String())
			}
		})
	}

}
