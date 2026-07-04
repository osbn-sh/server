package manipulationRepository

import (
	"ostadbun/entity"
	"ostadbun/pkg/richerror"
)

func (d DB) StabilizeUniversity(pendingUniversityID int) (submitterID int64, err error) {

	tx, err := d.conn.Conn().Beginx()
	if err != nil {
		return 0, richerror.New("manipulationRepository-StabilizeUniversity").WithErr(err).WithKind(richerror.KindUnexpected).WithMessage("error on begin transaction")
	}

	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			err = tx.Commit()
		}
	}()

	var pending entity.PendingUniversity

	fetchQuery := `
		SELECT * FROM pending_university
		WHERE status = 'approved' AND id = $1
		FOR UPDATE
	`

	err = tx.Get(&pending, fetchQuery, pendingUniversityID)
	if err != nil {
		return 0, richerror.New("manipulationRepository-StabilizeUniversity").WithErr(err).WithKind(richerror.KindUnexpected).WithMessage("error on get pending university")
	}

	insertQuery := `
		INSERT INTO university (
			name,
			name_english,
			city,
			category,
			image_url,
			description,
			description_english,
			registered_by
		)
		VALUES ($1,$2,$3,$4,$5,$6,$7,$8)
	`

	_, err = tx.Exec(
		insertQuery,
		pending.Name,
		pending.NameEnglish,
		pending.City,
		pending.Category,
		pending.ImageUrl,
		pending.Description,
		pending.DescriptionEnglish,
		pending.SubmittedBy,
	)
	if err != nil {
		return 0, richerror.New("manipulationRepository-StabilizeUniversity").WithErr(err).WithKind(richerror.KindUnexpected).WithMessage("error on insert pending university")
	}

	_, errE := tx.Exec(`DELETE FROM pending_university WHERE id = $1`, pendingUniversityID)

	if errE != nil {
		return 0, richerror.New("manipulationRepository-StabilizeUniversity").WithErr(err).WithKind(richerror.KindUnexpected).WithMessage("error on delete pending university")
	}

	return pending.SubmittedBy, nil
}
