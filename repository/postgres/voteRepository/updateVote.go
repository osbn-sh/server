package voteRepository

import (
	"ostadbun/pkg/richerror"
)

func (d DB) ChangeRate(rateID, rate int) error {

	query := `
		UPDATE vote
		SET rate = $1
		WHERE id = $2
	`

	res, err := d.conn.Conn().Exec(query, rate, rateID)
	if err != nil {
		return richerror.New("voteRepository-ChangeRate").WithErr(err).WithKind(richerror.KindUnexpected)
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return richerror.New("voteRepository-ChangeRate").WithErr(err).WithKind(richerror.KindUnexpected)
	}

	if rowsAffected == 0 {
		return richerror.New("voteRepository-ChangeRate").WithMessage("vote not found").WithKind(richerror.KindNotFound)
	}

	return nil
}
