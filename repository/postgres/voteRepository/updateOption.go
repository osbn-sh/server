package voteRepository

import (
	"ostadbun/param/voteparam"
	"ostadbun/pkg/richerror"
)

func (d DB) ChangeOption(optionID int, option voteparam.Option) error {

	query := `
		UPDATE option
		SET name = $1, weight = $2, owner = $3
		WHERE id = $4
	`

	res, err := d.conn.Conn().Exec(query, option.Name, option.Weight, option.Owner, optionID)
	if err != nil {
		return richerror.New("voteRepository-ChangeOption").WithErr(err).WithKind(richerror.KindUnexpected)
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return richerror.New("voteRepository-ChangeOption").WithErr(err).WithKind(richerror.KindUnexpected)
	}

	if rowsAffected == 0 {
		return richerror.New("voteRepository-ChangeOption").WithMessage("option not found").WithKind(richerror.KindNotFound)
	}

	return nil
}
