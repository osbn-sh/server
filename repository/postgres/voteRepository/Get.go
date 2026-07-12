package voteRepository

import (
	"ostadbun/pkg/richerror"
)

// ستون‌های مجاز برای جلوگیری از SQL Injection (چون اسم ستون رو نمی‌شه پارامتر کرد)
var allowedEntityColumns = map[string]string{
	"university": "university_id",
	"professor":  "professor_id",
	"major":      "major_id",
	"lesson":     "lesson_id",
}

func (d DB) CalcVotes(entityType string, entityID int) (float64, error) {

	column, ok := allowedEntityColumns[entityType]
	if !ok {
		return 0, richerror.New("voteRepository-CalcVotes").
			WithMessage("invalid entity type").
			WithKind(richerror.KindInvalid)
	}

	query := `
		SELECT
			COALESCE(SUM(v.rate * o.weight)::float / NULLIF(SUM(o.weight), 0), 0)
		FROM vote v
		JOIN option o ON o.id = v.option_id AND o.owner = $2
		WHERE v.` + column + ` = $1
	`

	var average float64
	err := d.conn.Conn().QueryRow(query, entityID, entityType).Scan(&average)
	if err != nil {
		return 0, richerror.New("voteRepository-CalcVotes").WithErr(err).WithKind(richerror.KindUnexpected)
	}

	return average, nil
}
