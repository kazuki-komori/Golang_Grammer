package example

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func Router() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", home)
	router.HandleFunc("/users", findAllUsers)
	log.Fatal(http.ListenAndServe(":8080", router))
}

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World")
}

type User struct {
	ID        int
	FirstName string
	LastName  string
}

func findAllUsers(w http.ResponseWriter, r *http.Request) {
	var userList = []User{
		{ID: 1, FirstName: "山田", LastName: "太郎"},
		{ID: 2, FirstName: "鈴木", LastName: "一郎"},
	}

	response, _ := json.Marshal(userList)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}
