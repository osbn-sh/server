package userRepository

import (
	"ostadbun/entity"
	"ostadbun/pkg/richerror"
)

func (a DB) GetUserByID(id int) (entity.User, error) {

	var userData entity.User

	query := `select id,university_id,major_id,admin_by from users where id = $1;`

	err := a.conn.Conn().QueryRow(query, id).Scan(
		&userData.Id,
		&userData.UniversityId,
		&userData.MajorId,
		&userData.AdminBy,
	)

	if err != nil {
		return entity.User{}, richerror.New("userRepository-GetUserByID").WithErr(err).WithKind(richerror.KindUnexpected)
	}

	return userData, nil

}
