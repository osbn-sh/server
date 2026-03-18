package userservice

import (
	"ostadbun/entity"
)

func (r User) GetByID(id int) (entity.User, error) {
	return r.repo.GetUserByID(id)
}
