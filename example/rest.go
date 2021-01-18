package example

import (
	"fmt"
	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"go_sample/db"
	"log"
	"net/http"
)

func Router() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", home)
	router.HandleFunc("/users", db.FindAllUsers).Methods("GET")
	router.HandleFunc("/users/{id}", db.FindById).Methods("GET")
	router.HandleFunc("/users", db.CreateUser).Methods("POST")
	log.Fatal(http.ListenAndServe(":8080", router))
}

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World")
}
