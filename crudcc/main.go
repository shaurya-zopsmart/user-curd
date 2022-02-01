package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	userHandler "github.com/shaurya-zopsmart/crudcc/http/user"
	Middlewares "github.com/shaurya-zopsmart/crudcc/middleware"
	userService "github.com/shaurya-zopsmart/crudcc/service/user"
	userStore "github.com/shaurya-zopsmart/crudcc/stores/users"
)

func connt() (*sql.DB, error) {
	conn := "sam:root@tcp(localhost:3306)/db1"
	db, err := sql.Open("mysql", conn)
	if err != nil {
		return nil, err
	}
	return db, nil
}

func main() {
	db, err := connt()
	if err != nil {
		log.Printf("error in db %v", err)
		return
	}
	err = db.Ping()
	if err != nil {
		log.Printf("error in connection %v", err)
		return
	}

	u := userStore.New(db)
	p := userService.New(&u)
	ht := userHandler.New(p)

	r := mux.NewRouter()

	r.Path("/user").Methods("POST").Handler(func() http.Handler {
		return Middlewares.Logger(http.HandlerFunc(ht.AddUser))
	}())
	// r.Path("/user").Methods("POST").HandlerFunc(ht.AddUser)
	r.Path("/user/{Id}").Methods("GET").HandlerFunc(ht.UserwithID)
	r.Path("/user").Methods("GET").HandlerFunc(ht.GetAllUsers)
	r.Path("/user/{id}").Methods("PUT").HandlerFunc(ht.UpdateUser)
	// r.Path("/user",midd.Logger(http.HandlerFunc(ht.UpdateUser)).Methods(http.MethodPut)
	r.Path("/user/{id}").Methods("DELETE").HandlerFunc(ht.DeleteUser)
	http.Handle("/", r)
	fmt.Println("listing tp port 3000")
	err = http.ListenAndServe(":3000", nil)
	if err != nil {
		fmt.Println(err)
	}
}
