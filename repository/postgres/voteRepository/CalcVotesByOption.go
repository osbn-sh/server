package voteRepository

import (
	"ostadbun/entity"
	"ostadbun/pkg/richerror"
)

func (d DB) CalcVotesByOption(entityType string, entityID int) ([]entity.OptionVoteResult, error) {

	column, ok := allowedEntityColumns[entityType]
	if !ok {
		return nil, richerror.New("voteRepository-CalcVotesByOption").
			WithMessage("invalid entity type").
			WithKind(richerror.KindInvalid)
	}

	query := `
		SELECT
			o.id,
			o.name,
			o.weight,
			COALESCE(AVG(v.rate)::float, 0) AS average_rate,
			COUNT(v.rate) AS vote_count
		FROM option o
		LEFT JOIN vote v ON v.option_id = o.id AND v.` + column + ` = $1
		WHERE o.owner = $2
		GROUP BY o.id, o.name, o.weight
		ORDER BY o.id
	`

	rows, err := d.conn.Conn().Query(query, entityID, entityType)
	if err != nil {
		return []entity.OptionVoteResult{}, richerror.New("voteRepository-CalcVotesByOption").WithErr(err).WithKind(richerror.KindUnexpected)
	}
	defer rows.Close()

	results := make([]entity.OptionVoteResult, 0)
	for rows.Next() {
		var r entity.OptionVoteResult
		if err := rows.Scan(&r.OptionID, &r.OptionName, &r.Weight, &r.AverageRate, &r.VoteCount); err != nil {
			return nil, richerror.New("voteRepository-CalcVotesByOption").WithErr(err).WithKind(richerror.KindUnexpected)
		}
		results = append(results, r)
	}

	if err := rows.Err(); err != nil {
		return nil, richerror.New("voteRepository-CalcVotesByOption").WithErr(err).WithKind(richerror.KindUnexpected)
	}

	return results, nil
}
