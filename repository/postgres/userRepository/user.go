package userRepository

import (
	"fmt"
	"ostadbun/entity"
	"ostadbun/pkg/hash"
)

func (a DB) ExistingCheck(email string) (int, bool) {
	var id int

	email = hash.Hasher(email)

	err := a.conn.Conn().QueryRow(
		`SELECT id FROM users WHERE email = $1 LIMIT 1`,
		email,
	).Scan(&id)

	fmt.Println(email, id)
	if err != nil {
		fmt.Println(err.Error())
		return 0, false

	}

	return id, id > 0
}

func (a DB) RegisterUser(user entity.User) (int, error) {

	var id int

	err := a.conn.Conn().QueryRow(`
		INSERT INTO users (email)
		VALUES ($1)
		RETURNING id
	`,
		user.Email_Hashe(),
	).Scan(&id)

	if err != nil {
		return 0, err
	}

	return id, nil
}
