package usecase

import "github.com/kuropenguin/go-clean-arch-copy/entity"

type UserRepository interface {
	Insert(user *entity.User) error
}
