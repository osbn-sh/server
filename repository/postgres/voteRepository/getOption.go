package voteRepository

import (
	"ostadbun/param/voteparam"
	"ostadbun/pkg/richerror"
)

func (d DB) GetOption() ([]voteparam.Option, error) {
	query := `
		SELECT id, name, weight, owner
		FROM option
	`

	rows, err := d.conn.Conn().Query(query)
	if err != nil {
		return nil, richerror.New("voteRepository-getOption").WithErr(err)
	}
	defer rows.Close()

	var options []voteparam.Option

	for rows.Next() {
		var option voteparam.Option

		err := rows.Scan(
			&option.OptionID,
			&option.Name,
			&option.Weight,
			&option.Owner,
		)
		if err != nil {
			return nil, richerror.New("voteRepository-getOption-scan").WithErr(err)
		}

		options = append(options, option)
	}

	if err := rows.Err(); err != nil {
		return nil, richerror.New("voteRepository-getOption").WithErr(err)
	}

	return options, nil
}
