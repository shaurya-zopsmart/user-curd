package user

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/shaurya-zopsmart/user-curd/crudcc/models"
	"github.com/shaurya-zopsmart/user-curd/crudcc/service"
)

type Handler struct {
	handler service.User
}

func New(s service.User) *Handler {
	return &Handler{
		handler: s,
	}
}

func (h Handler) UserwithID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	para := mux.Vars(r)
	v := para["Id"]
	id, err := strconv.Atoi(v)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		newerror := models.ErrorResponse{StatusCode: http.StatusBadRequest, ErrorMessage: "Bad request , User id invalid"}
		err, _ := json.Marshal(newerror)
		_, _ = w.Write(err)
		return
	}

	user, err := h.handler.GetUsrById(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		newError := models.ErrorResponse{StatusCode: http.StatusBadRequest, ErrorMessage: "BAD request, User id not found"}
		err, _ := json.Marshal(newError)
		_, _ = w.Write(err)
		return
	}

	w.WriteHeader(http.StatusOK)
	resp := models.Response{Data: user, Message: "User retrived successfully ", StatusCode: 200}
	data, _ := json.Marshal(resp)
	_, _ = w.Write(data)

}

func (us Handler) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	userdata, err := us.handler.GetAllUsrs()

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("no data found"))
		return
	}

	res, err := json.Marshal(userdata)
	if err == nil {
		w.Write([]byte(res))
		w.WriteHeader(http.StatusOK)
	}

}

func (us Handler) Create(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var ur models.User
	body := r.Body
	err := json.NewDecoder(body).Decode(&ur)

	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("invalid body"))
	}

	res, newerr := us.handler.Create(ur)
	if newerr != nil {
		fmt.Println(newerr)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(" c block"))
	}
	nres, err := json.Marshal(res)
	if err == nil {
		w.WriteHeader(http.StatusCreated)
		w.Write(nres)
	}

}

func (us Handler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var ur models.User
	body := r.Body
	json.NewDecoder(body).Decode(&ur)
	params := mux.Vars(r)
	v := params["id"]
	id, _ := strconv.Atoi(v)

	err := us.handler.UpdateUsr(id, ur)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("cannot update due to some validation error"))
		return
	}
	b, _ := json.Marshal(ur)
	w.WriteHeader(http.StatusAccepted)
	w.Write(b)

}

func (us Handler) DeleteUser(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	v := params["id"]
	id, _ := strconv.Atoi(v)

	err := us.handler.DeleteUsrById(id)
	if err != nil {

		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("User Deletion Failed"))
		return
	}
	w.WriteHeader(http.StatusAccepted)
	w.Write([]byte("User Deleted"))

}
