package db

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
)

type User struct {
	ID        int
	FirstName string
	LastName  string
}

func connectToDB() *gorm.DB {
	err := godotenv.Load(fmt.Sprintf("./%s.env", os.Getenv("GO_ENV")))
	env := os.Getenv("ENV")
	DB := os.Getenv("DB")
	DBUser := os.Getenv("DB_USER")
	DBPass := os.Getenv("DB_PASS")

	CONNECT := DBUser + ":" + DBPass + "@" + env + "/" + DB
	fmt.Println(CONNECT)

	db, err := gorm.Open("mysql", CONNECT)

	if err != nil {
		panic(err.Error())
	}
	return db
}

func FindAllUsers(w http.ResponseWriter, r *http.Request) {
	db := connectToDB()
	defer db.Close()

	db.LogMode(true)

	var userList []User

	db.Find(&userList)
	fmt.Println(userList)
	fmt.Println(&userList)

	response, _ := json.Marshal(userList)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

func FindById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])

	if err != nil {
		log.Fatal(err)
	}

	db := connectToDB()
	defer db.Close()

	db.LogMode(true)

	var user User

	db.Where("id = ?", id).Find(&user)

	response, _ := json.Marshal(user)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		log.Fatal(err)
	}
	var user User
	json.Unmarshal(body, &user)
	db := connectToDB()
	defer db.Close()

	db.Create(&user)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	response, _ := json.Marshal(user)
	w.Write(response)
}
