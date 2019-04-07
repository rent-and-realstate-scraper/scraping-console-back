package models

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	"os"
)



var dbGorm *gorm.DB
var db *sqlx.DB

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
	db, _ = sqlx.Connect("mysql",  os.Getenv("database_url"))

	sqlCreation := "create table if not exists accounts (ID int key NOT NULL AUTO_INCREMENT, token varchar(20), password varchar(20), email varchar (20));"
	_, errCr := db.Queryx(sqlCreation)
	if errCr!= nil{
		panic(err)
	}
}

func GetDBGorm() *gorm.DB {
	return dbGorm
}

func GetDb() *sqlx.DB {
	return db
}
