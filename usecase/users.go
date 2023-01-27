package usecase

import "github.com/kuropenguin/go-clean-arch-copy/entity"

type UserUsecase interface {
	CreateUser(name string, age int) error
}

type userUsecaseImpl struct {
	userRepository UserRepository
}

func (u *userUsecaseImpl) CreateUser(name string, age int) error {
	user, err := entity.NewUser(name, age)
	if err != nil {
		return err
	}

	err = u.userRepository.Insert(user)
	if err != nil {
		return err
	}

	return nil
}
