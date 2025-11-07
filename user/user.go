package user

import (
	"errors"
	"fmt"
	"time"
)

type User struct {
	firstName string
	lastName  string
	birthDate string
	createdAt time.Time
}

type Admin struct {
	email    string
	password string
	User
}

func (u *User) OutputUserDetails() {
	fmt.Println(u.firstName, u.lastName, u.birthDate)
}

func (admin Admin) OutputUserDetails() {
	fmt.Println(admin.email, admin.password, admin.firstName, admin.lastName, admin.birthDate)
}

func (u *User) ClearUserName() {
	u.firstName = "-"
	u.lastName = "-"
}

func New(firstName, lastName, birthDate string) (*User, error) {
	if firstName == "" || lastName == "" || birthDate == "" {
		return nil, errors.New("Error: First name, last name and birth date cannot be empty")
	}

	return &User{
		firstName: firstName,
		lastName:  lastName,
		birthDate: birthDate,
		createdAt: time.Now(),
	}, nil
}

func NewAdmin(email, password, firstName, lastName, birthDate string) (*Admin, error) {
	if email == "" || password == "" || firstName == "" || lastName == "" || birthDate == "" {
		return nil, errors.New("Error: All fields cannot be empty")
	}

	return &Admin{
		email:    email,
		password: password,
		User: User{
			firstName: firstName,
			lastName:  lastName,
			birthDate: birthDate,
			createdAt: time.Now(),
		},
	}, nil
}
