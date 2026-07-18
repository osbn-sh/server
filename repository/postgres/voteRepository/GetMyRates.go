package voteRepository

import (
	"ostadbun/entity"
	"ostadbun/pkg/richerror"
)

func (d DB) GetMyRates(entityType string, userID, targetID int) ([]entity.MyVote, error) {
	query := `
		SELECT
			v.id,
			v.option_id,
			v.rate,
			o.name
		FROM vote AS v
		INNER JOIN option AS o
			ON o.id = v.option_id
		WHERE v.user_id = $1
	`

	args := []any{userID}

	switch entityType {
	case "university":
		query += " AND v.university_id = $2"
		args = append(args, targetID)

	case "professor":
		query += " AND v.professor_id = $2"
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
			&v.OptionName,
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
