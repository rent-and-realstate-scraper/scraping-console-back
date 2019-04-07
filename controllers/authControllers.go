package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"os"
	"scraping-console-back/models"
	u "scraping-console-back/utils"
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

	resp, code:= Login(account.Email, account.Password)
	if code == 401 {
		w.WriteHeader(http.StatusUnauthorized)
	} else if code == 500 {
		w.WriteHeader(http.StatusInternalServerError)
	}
	u.Respond(w, resp)
}

func Login(email, password string) (response map[string]interface{}, code int) {

	account := &models.Account{}
	err := models.GetDBGorm().Table("accounts").Where("email = ?", email).First(account).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return u.Message(false, "Email address not found"), 401
		}
		return u.Message(false, "Connection error. Please retry"), 500
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
	models.GetDBGorm().Table("accounts").Where("id = ?", u).First(acc)
	if acc.Email == "" { //User not found!
		return nil
	}

	acc.Password = ""
	return acc
}
