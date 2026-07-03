package userRepository

import (
	"ostadbun/param/userparam"
	"ostadbun/pkg/richerror"
	"strings"
)

func (a DB) RegisterUserByEmailAndPassword(user userparam.User) error {

	query := `INSERT into "users" (email,password) values ($1,$2)`

	err := a.conn.Conn().QueryRow(query, user.Email, user.Password).Err()

	if err != nil {

		if strings.Contains(err.Error(), "23505") {
			return richerror.New("userRepository-RegisterUserByEmailAndPassword").WithErr(err).WithKind(richerror.KindUnexpected)
		}

		return richerror.New("userRepository-RegisterUserByEmailAndPassword").WithErr(err).WithKind(richerror.KindUnexpected)
	}
	return nil
}
