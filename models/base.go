package models

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
	"os"
)



var dbGorm *gorm.DB
var db *sql.DB

func init() {

	e := godotenv.Load()
	if e != nil {
		fmt.Print(e)
	}

	username := os.Getenv("db_user")
	password := os.Getenv("db_pass")
	dbName := os.Getenv("db_name")
	dbHost := os.Getenv("db_host")
	dbAddress := os.Getenv("database_url")



	dbUri := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s", dbHost, username, dbName, password)
	fmt.Println(dbUri)

	conn, err := gorm.Open("mysql", dbAddress)
	if err != nil {
		fmt.Println("error")
		fmt.Print(err)
	}

	dbGorm = conn
	dbGorm.Debug().AutoMigrate(&Account{}, &Contact{})
	db, _ = sql.Open("mysql",  os.Getenv("database_url"))
}

func GetDBGorm() *gorm.DB {
	return dbGorm
}

func GetDb() *sql.DB {
	return db
}
