package userservice

import (
	"fmt"
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

	fmt.Println("enterd into repo layer", newHashedData)
	a, b := r.repo.GetUserByEmailAndPass(newHashedData)

	fmt.Println("repo layer", a, b)
	return a, b

}
