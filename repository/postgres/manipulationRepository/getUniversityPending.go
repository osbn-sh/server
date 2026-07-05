package manipulationRepository

import (
	"database/sql"
	"ostadbun/entity"
	"ostadbun/pkg/richerror"
)

func (d DB) GetUniversityPending(Id int, filterBySubmitter bool) ([]entity.PendingUniversity, error) {
	query := `
        SELECT 
            id,
            name,
            name_english,
            description_english,
            city,
            category,
            image_url,
            description,
            status,
            submitted_by,
            submitted_at,
            approved_by,
            approved_at,
            rejection_reason,
            action,
            target_id
        FROM pending_university
        WHERE status = 'pending'
    `

	var (
		rows *sql.Rows
		err  error
	)

	if Id > 0 {
		if filterBySubmitter {
			query += " AND submitted_by = $1"
		} else {
			query += " AND id = $1"
		}
		rows, err = d.conn.Conn().Query(query, Id)
	} else {
		rows, err = d.conn.Conn().Query(query)
	}

	if err != nil {
		return nil, richerror.New("manipulationRepository-GetUniversityPending").WithErr(err).WithKind(richerror.KindUnexpected).WithMessage("error on query pending university")
	}
	defer rows.Close()

	var universities []entity.PendingUniversity

	for rows.Next() {
		var university entity.PendingUniversity
		err := rows.Scan(
			&university.Id,
			&university.Name,
			&university.NameEnglish,
			&university.DescriptionEnglish,
			&university.City,
			&university.Category,
			&university.ImageUrl,
			&university.Description,
			&university.Status,
			&university.SubmittedBy,
			&university.SubmittedAt,
			&university.ApprovedBy,
			&university.ApprovedAt,
			&university.RejectionReason,
			&university.Action,
			&university.TargetId,
		)
		if err != nil {
			return nil, richerror.New("manipulationRepository-GetUniversityPending").WithErr(err).WithKind(richerror.KindUnexpected).WithMessage("error on query pending university")
		}
		universities = append(universities, university)
	}

	if err := rows.Err(); err != nil {
		return nil, richerror.New("manipulationRepository-GetUniversityPending").WithErr(err).WithKind(richerror.KindUnexpected).WithMessage("error on query pending university")
	}

	return universities, nil
}
