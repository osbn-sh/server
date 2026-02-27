package userRepository

import "ostadbun/pkg/richerror"

func (a DB) AdminByWho(userID string) (int, error) {

	var userIdFromDB int

	err := a.conn.Conn().QueryRow("select admin_by from users where id=$1", userID).Scan(&userIdFromDB)

	if err != nil {
		return 0, richerror.New("userRepository-AdminByWho").WithErr(err).WithMessage("").WithKind(richerror.KindUnexpected)
	}

	return userIdFromDB, nil
}
