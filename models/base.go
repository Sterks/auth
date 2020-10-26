package models

import (
	"fmt"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
)

var db *gorm.DB

func init() {

	e := godotenv.Load()
	if e != nil {
		fmt.Print(e)
	}

	username := os.Getenv("db_user")
	password := os.Getenv("db_pass")
	dbName := os.Getenv("db_name")
	dbHost := os.Getenv("db_host")

	dbUri := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s", dbHost, username, dbName, password)
	fmt.Println(dbUri)

	type Config struct {
		connect string
	}

	var config Config
	config.connect = dbUri

	conn, err := gorm.Open(postgres.Open(dbUri), &gorm.Config{})

		if err != nil {
		fmt.Print(err)
	}

	db = conn
	err2 := db.Debug().AutoMigrate(&Account{}, &Role{}, &Contact{})
	if err2 != nil {
		log.Println(err2)
	}
}

func GetDB() *gorm.DB {
	return db
}
