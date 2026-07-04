package manipulationRepository

import (
	"database/sql"
	"ostadbun/entity"
	"ostadbun/pkg/richerror"
)

// GetMajorPending returns all majors with 'pending' status
func (d DB) GetMajorPending(userId int) ([]entity.PendingMajor, error) {
	query := `
	SELECT
		id,
		name,
		status,
		name_english,
		submitted_by,
		description,
		submitted_at,
		description_english,
		approved_by,
		approved_at,
		rejection_reason
	FROM pending_major
	WHERE status = 'pending'
`

	var (
		rows *sql.Rows
		err  error
	)

	if userId > 0 {
		query += " AND submitted_by = $1"
		rows, err = d.conn.Conn().Query(query, userId)
	} else {
		rows, err = d.conn.Conn().Query(query)
	}

	//rows, err := d.conn.Conn().Query(query, userId)
	if err != nil {
		return nil, richerror.New("manipulationRepository-GetMajorPending").WithErr(err).WithKind(richerror.KindUnexpected).WithMessage("error on query pending major")
	}
	defer rows.Close()

	var majors []entity.PendingMajor

	for rows.Next() {
		var major entity.PendingMajor
		err := rows.Scan(
			&major.Id,
			&major.Name,
			&major.Status,
			&major.NameEnglish,
			&major.SubmittedBy,
			&major.Description,
			&major.SubmittedAt,
			&major.DescriptionEnglish,
			&major.ApprovedBy,
			&major.ApprovedAt,
			&major.RejectionReason,
		)
		if err != nil {
			return nil, richerror.New("manipulationRepository-GetMajorPending").WithErr(err).WithKind(richerror.KindUnexpected).WithMessage("error on query pending major")
		}
		majors = append(majors, major)
	}

	if err := rows.Err(); err != nil {
		return nil, richerror.New("manipulationRepository-GetMajorPending").WithErr(err).WithKind(richerror.KindUnexpected).WithMessage("error on query pending major")
	}

	return majors, nil
}
