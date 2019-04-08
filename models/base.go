package models

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	"os"
)

var db *sqlx.DB

func init() {

	e := godotenv.Load()
	if e != nil {
		fmt.Print(e)
	}

 	db, _ = sqlx.Connect("mysql",  os.Getenv("database_url"))

	sqlCreationAccount := "create table if not exists accounts (id int key NOT NULL AUTO_INCREMENT, token varchar(20), password varchar(20), email varchar (20));"
	sqlCreationContacts := "create table if not exists contacts (id int key NOT NULL AUTO_INCREMENT, user_id int, name varchar(20), phone varchar (20));"

	_, errCr1 := db.Queryx(sqlCreationAccount)
	_, errCr2 := db.Queryx(sqlCreationContacts)

	if errCr1!= nil || errCr2!= nil{
		panic(errCr1)
	}
}

func GetDb() *sqlx.DB {
	return db
}
