package voteRepository

import (
	"ostadbun/entity"
	"ostadbun/pkg/richerror"
)

func (d DB) GetMyRates(entityType string, userID, targetID int) ([]entity.MyVote, error) {
	query := `
		SELECT
			id,
			option_id,
			rate
		FROM vote
		WHERE user_id = $1
	`

	args := []any{userID}

	switch entityType {
	case "university":
		query += " AND university_id = $2"
		args = append(args, targetID)

	case "professor":
		query += " AND professor_id = $2"
		args = append(args, targetID)

	default:
		return nil, richerror.New("voteRepository-GetMyRates").
			WithMessage("invalid entity type").
			WithKind(richerror.KindInvalid)
	}

	rows, err := d.conn.Conn().Query(query, args...)
	if err != nil {
		return nil, richerror.New("voteRepository-GetMyRates").
			WithErr(err).
			WithKind(richerror.KindUnexpected)
	}
	defer rows.Close()

	results := make([]entity.MyVote, 0)

	for rows.Next() {
		var v entity.MyVote

		if err := rows.Scan(
			&v.Id,
			&v.OptionId,
			&v.Rate,
		); err != nil {
			return nil, richerror.New("voteRepository-GetMyRates").
				WithErr(err).
				WithKind(richerror.KindUnexpected)
		}

		results = append(results, v)
	}

	if err := rows.Err(); err != nil {
		return nil, richerror.New("voteRepository-GetMyRates").
			WithErr(err).
			WithKind(richerror.KindUnexpected)
	}

	return results, nil
}
