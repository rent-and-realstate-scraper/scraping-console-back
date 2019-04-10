package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"scraping-console-back/models"
	u "scraping-console-back/utils"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

var CreateAccount = func(w http.ResponseWriter, r *http.Request) {

	account := &models.Account{}
	err := json.NewDecoder(r.Body).Decode(account) //decode the request body into struct and failed if any error occur
	if err != nil {
		u.Respond(w, u.Message(false, "Invalid request"))
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	fmt.Println(*account)
	resp := account.Create() //Create account
	u.Respond(w, resp)
}

var Authenticate = func(w http.ResponseWriter, r *http.Request) {

	account := &models.Account{}
	err := json.NewDecoder(r.Body).Decode(account) //decode the request body into struct and failed if any error occur
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		u.Respond(w, u.Message(false, "Invalid request"))
		return
	}

	resp, code := Login(account.Email, account.Password)
	if code == 401 {
		w.WriteHeader(http.StatusUnauthorized)
	} else if code == 500 {
		w.WriteHeader(http.StatusInternalServerError)
	}
	u.Respond(w, resp)
}

func Login(email, password string) (response map[string]interface{}, code int) {

	var account models.Account

	db := models.GetDb()
	sql := fmt.Sprintf("select * from accounts where email = '%s'", email)

	rows := db.QueryRowx(sql)
	err := rows.StructScan(&account)

	if err != nil {
		fmt.Println(err)
		return u.Message(false, "Email address not found"), 401
	}

	err = bcrypt.CompareHashAndPassword([]byte(account.Password), []byte(password))
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword { //Password does not match!
		return u.Message(false, "Invalid login credentials. Please try again"), 401
	}
	//Worked! Logged In
	account.Password = ""

	//Create JWT token
	tk := &models.Token{UserId: account.ID}
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)
	tokenString, _ := token.SignedString([]byte(os.Getenv("token_password")))
	account.Token = tokenString //Store the token in the response

	resp := u.Message(true, "Logged In")
	resp["account"] = account
	return resp, 200
}

func GetUser(u uint) *models.Account {

	acc := &models.Account{}

	db := models.GetDb()
	sql := fmt.Sprintf("select * from accounts where id = '%d'", u)

	rows := db.QueryRowx(sql)
	rows.StructScan(&acc)

	if acc.Email == "" {
		return nil
	}

	acc.Password = ""
	return acc
}
