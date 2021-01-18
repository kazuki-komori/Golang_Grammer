package example

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"

	_ "github.com/jinzhu/gorm/dialects/mysql"
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
