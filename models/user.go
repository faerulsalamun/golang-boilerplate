package models

import (
	"github.com/faerulsalamun/golang-boilerplate/db"
)

type User struct {
	Name string `json:"name"`
}

func (h User) GetAll() ([] User, error) {

	con := db.Init().C("user")

	defer db.CloseSession()

	var result [] User

	err := con.Find(nil).All(&result)

	if err != nil {
		return nil, err
	}

	return result, nil
}