package userservice

import (
	"fmt"
	"ostadbun/entity"
	"ostadbun/param/userparam"
)

func (r User) GetByID(id int) (entity.User, error) {
	return r.repo.GetUserByID(id)
}

func (r User) GetByEmailAndPass(email, pass string) (*int64, error) {
	user := userparam.User{
		Email:    email,
		Password: pass,
	}

	fmt.Println("enterd into repo layer")
	a, b := r.repo.GetUserByEmailAndPass(user)

	fmt.Println("repo layer", a, b)
	return a, b

}
