package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// tutorial values
var user = "prometheus"
var password = "Password1"
var host = "youtubetesting.mysql.database.azure.com"
var port = "3306"
var database = "go_test_models"

// he was using mysql hosted on azure, I'll be doing aws, and developing with a local dockerized instance of mysql this might change some things

// Database Connection
var dsn = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4", user, password, host, port, database) // should the tcp be parameterized?
var db, _ = gorm.Open(mysql.Open(dsn), &gorm.Config{})

type GoTestModel struct {
	Name string
	Year string
}

func GoDatabaseCreate(w http.ResponseWriter, r *http.Request) {
	GoTestModel := GoTestModel{
		Name: "Mike",
		Year: "2021",
	}

	db.Create(&GoTestModel)
	if err := db.Create(&GoTestModel).Error; err != nil {
		log.Fatalln((err))
	}

	json.NewEncoder(w).Encode(GoTestModel)

	fmt.Println("Fields Added", GoTestModel)
}

func main() {
	// we need introduce more robust concurrency mgmt
	run()
}

func run() {
	http.HandleFunc("/createstuff", GoDatabaseCreate)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
