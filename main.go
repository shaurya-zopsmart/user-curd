package main

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	userHandler "github.com/shaurya-zopsmart/crudcc/http/user"
	userService "github.com/shaurya-zopsmart/crudcc/service/user"
	userStore "github.com/shaurya-zopsmart/crudcc/stores/users"
)

func main() {
	config := mysql.Config{
		User:   "sam",
		Passwd: "root",
		Net:    "tcp",
		Addr:   "localhost:3306",
		DBName: "user",
	}
	var err error
	db, err := sql.Open("mysql", config.FormatDSN())
	if err != nil {
		fmt.Println(err)
		fmt.Println("Error in connection establishment!")
		return
	}
	u := userStore.New(db)
	p := userService.New(&u)
	ht := userHandler.Handler{p}
	r := mux.NewRouter().StrictSlash(true)
	r.HandleFunc("/user/{id}", ht.UserwithID).Methods(http.MethodGet)
	r.HandleFunc("/user", ht.AddUser).Methods(http.MethodPost)
	r.HandleFunc("/user", ht.GetAllUsers).Methods(http.MethodGet)
	r.HandleFunc("/user/delete/{id}", ht.DeleteUser).Methods(http.MethodDelete)
	r.HandleFunc("/user/update/{id}", ht.UpdateUser).Methods(http.MethodPatch)
	http.ListenAndServe(":8080", r)
}
