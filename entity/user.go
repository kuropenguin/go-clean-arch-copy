package entity

import "errors"

type User struct {
	name string
	age  int
}

func NewUser(name string, age int) (*User, error) {
	if age < 20 {
		return nil, errors.New("too young")
	}
	return &User{
		name: name,
		age:  age,
	}, nil
}

func (u *User) Name() string {
	return u.name
}

func (u *User) Age() int {
	return u.age
}
