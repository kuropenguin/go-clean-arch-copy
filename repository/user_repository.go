package repository

import (
	"github.com/kuropenguin/go-clean-arch-copy/entity"
)

type userRepository struct{}

type myDB struct {
	lastInsertID int
	users        map[int]entity.User
}

var db myDB

func (r *userRepository) Insert(user *entity.User) error {
	db.lastInsertID++
	db.users[db.lastInsertID] = *user
	return nil
}
