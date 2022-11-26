package db

import (
	"log"

	"github.com/Max-Gabriel-Susman/bestir-identity-service/internal/account"
	"github.com/jinzhu/gorm"
)

func InitializeDB() error {
	db, err := gorm.Open("mysql", "usr:identity@tcp(0.0.0.0:3306)/identity?multiStatements=true&parseTime=true")
	if err != nil {
		log.Println("Connection failed to open")
		return err
	}
	defer db.Close()
	log.Println("Connection established")

	db.CreateTable(&account.Account{})

	return nil
}
