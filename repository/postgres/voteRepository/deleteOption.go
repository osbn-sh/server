package voteRepository

import (
	"ostadbun/pkg/richerror"
)

func (d DB) RemoveOption(id int) error {

	query := `
		DELETE FROM option WHERE id = $1
	`

	res, err := d.conn.Conn().Exec(query, id)
	if err != nil {
		return richerror.New("voteRepository-RemoveOption").WithErr(err).WithKind(richerror.KindUnexpected)
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return richerror.New("voteRepository-RemoveOption").WithErr(err).WithKind(richerror.KindUnexpected)
	}

	if rowsAffected == 0 {
		return richerror.New("voteRepository-RemoveOption").WithMessage("option not found").WithKind(richerror.KindNotFound)
	}

	return nil
}
