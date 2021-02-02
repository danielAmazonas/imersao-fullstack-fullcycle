package desafios

import (
	uuid "github.com/satori/go.uuid"
)

type User struct {
	ID string `json:"id" valid:"uuid"`
	Name string `json:"name" valid:"notnull"`
	Email string `json:"email" valid:"notnull"`
}

func NewUser(name string, email string) (*User, error) {
	user := User{
		Name: name,
		Email: email,
	}
	user.ID = uuid.NewV4().String()
	return &user, nil
}