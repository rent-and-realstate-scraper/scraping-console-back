package models

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
	"os"
	u "scraping-console-back/utils"
	"strings"
)

type Token struct {
	UserId uint
	jwt.StandardClaims
}

type Account struct {
	Email string `json:"email"`
	Password string `json:"password"`
	Token string `json:"token"`
	ID uint `json:"id"`
}

func (account *Account) Validate() (map[string] interface{}, bool) {

	if !strings.Contains(account.Email, "@") {
		return u.Message(false, "Email address is required"), false
	}

	if len(account.Password) < 6 {
		return u.Message(false, "Password is required"), false
	}

	//Email must be unique
	temp := Account{}

	db = GetDb()
	sql := fmt.Sprintf("select * from accounts where email = '%s'",account.Email)

	rows := db.QueryRowx(sql)

	err := rows.StructScan(&temp)
	if err != nil {
		fmt.Println(err)
		// return u.Message(false, "Connection error. Please retry"), false
	}


	if temp.Email != "" {
		return u.Message(false, "Email address already in use by another user."), false
	}

	return u.Message(false, "Requirement passed"), true
}

func (account *Account) Create() (map[string] interface{}) {

	if resp, ok := account.Validate(); !ok {
		return resp
	}

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(account.Password), bcrypt.DefaultCost)
	account.Password = string(hashedPassword)

	db := GetDb()
	fmt.Println(account.Token)
	sql :="insert into accounts (token, password, email) values ('"+ account.Token + "','"+  account.Password + "','"+account.Email + "')"
	fmt.Println(sql)
	_, err := db.Queryx(sql)

	if err != nil {
		fmt.Println(err)
		return u.Message(false, "Failed to create account, connection error.")
	}

	//Create new JWT token for the newly registered account
	tk := &Token{UserId: account.ID}
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)
	tokenString, _ := token.SignedString([]byte(os.Getenv("token_password")))
	account.Token = tokenString

	account.Password = "" //delete password

	response := u.Message(true, "Account has been created")
	response["account"] = account
	return response
}

