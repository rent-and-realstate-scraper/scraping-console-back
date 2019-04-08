package models

import (
	"fmt"
	u "scraping-console-back/utils"
)

type Contact struct {
	ID uint `json:"id"`
	Name string `json:"name"`
	Phone string `json:"phone"`
	UserId uint `json:"user_id"` //The user that this contact belongs to
}

/*
 This struct function validate the required parameters sent through the http request body

returns message and true if the requirement is met
*/
func (contact *Contact) Validate() (map[string] interface{}, bool) {

	if contact.Name == "" {
		return u.Message(false, "Contact name should be on the payload"), false
	}

	if contact.Phone == "" {
		return u.Message(false, "Phone number should be on the payload"), false
	}

	if contact.UserId <= 0 {
		return u.Message(false, "User is not recognized"), false
	}

	//All the required parameters are present
	return u.Message(true, "success"), true
}

func (contact *Contact) Create() (map[string] interface{}) {

	if resp, ok := contact.Validate(); !ok {
		return resp
	}

	sql :="insert into contacts (name, phone, user_id) values ('"+ contact.Name + "','"+  contact.Phone + "','" + string(contact.UserId) + "')"
	fmt.Println(sql)
	_, err := db.Queryx(sql)

	var resp map[string]interface{}
	if (err != nil) {
		fmt.Println(err)
		resp = u.Message(false, "error creating account")

	}
	resp = u.Message(true, "success")
	resp["contact"] = contact
	return resp
}

func GetContact(id uint) (*Contact) {

	contact := &Contact{}

	db = GetDb()
	sql := fmt.Sprintf("select * from contacts where id = '%d'",contact.ID)

	rows := db.QueryRowx(sql)
	err := rows.StructScan(&contact)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	if err != nil {
		return nil
	}
	return contact
}

func GetContacts(user uint) ([]Contact) {

	var contacts []Contact
	db = GetDb()
	sql := fmt.Sprintf("select * from contacts where id = '%d'",user)

	rows, err := db.Queryx(sql)
	for rows.Next() {
		var contact Contact
		err = rows.StructScan(&contact)
		if err != nil {
			fmt.Println(err)
		}
		contacts = append(contacts, contact)
	}

	return contacts
}

