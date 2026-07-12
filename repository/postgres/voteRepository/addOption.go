package voteRepository

import (
	"fmt"
	"ostadbun/param/voteparam"
	"ostadbun/pkg/richerror"
)

func (d DB) AddOption(option voteparam.Option) error {

	query := `
       	INSERT INTO option (name ,weight,owner) 
		VALUES ($1, $2,$3)
    `

	res, errT := d.conn.Conn().Exec(query,
		option.Name,
		option.Weight,
		option.Owner,
	)

	if errT != nil {
		return richerror.New("voteRepository-addoption").WithErr(errT)
	}
	fmt.Println(res.RowsAffected())
	return nil
}
