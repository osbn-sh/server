package userRepository

import (
	"ostadbun/param/userparam"
	"ostadbun/pkg/richerror"
)

func (a DB) GetUserByEmailAndPass(user userparam.User) (*int64, error) {

	var userData int64

	query := `select id from "users" where email = $1 and password = $2;`

	err := a.conn.Conn().QueryRow(query, user.Email, user.Password).Scan(&userData)

	if err != nil {
		return nil, richerror.New("userRepository-GetUserByEmailAndPass").WithErr(err).WithKind(richerror.KindUnexpected)
	}

	return &userData, nil

}
