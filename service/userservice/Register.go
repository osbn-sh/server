package userservice

import (
	"ostadbun/param/userparam"
	"ostadbun/pkg/hash"
)

func (r User) RegisterUserByEmailAndPassword(u userparam.User) error {

	//TODO hash data

	email := hash.Hasher(u.Email)

	password := hash.Hasher(u.Password)

	newHashedData := userparam.User{
		Email:    email,
		Password: password,
	}

	return r.repo.RegisterUserByEmailAndPassword(newHashedData)
}

func (r User) IsExist(u userparam.User) (*int64, error) {

	email := hash.Hasher(u.Email)

	password := hash.Hasher(u.Password)

	newHashedData := userparam.User{
		Email:    email,
		Password: password,
	}
	return r.repo.GetUserByEmailAndPass(newHashedData)

}
