package voteRepository

import (
	"ostadbun/pkg/richerror"
)

func (d DB) RemoveVote(id int) error {

	query := `
		DELETE FROM vote WHERE id = $1
	`

	res, err := d.conn.Conn().Exec(query, id)
	if err != nil {
		return richerror.New("voteRepository-RemoveVote").WithErr(err).WithKind(richerror.KindUnexpected)
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return richerror.New("voteRepository-RemoveVote").WithErr(err).WithKind(richerror.KindUnexpected)
	}

	if rowsAffected == 0 {
		return richerror.New("voteRepository-RemoveVote").WithMessage("vote not found").WithKind(richerror.KindNotFound)
	}

	return nil
}
