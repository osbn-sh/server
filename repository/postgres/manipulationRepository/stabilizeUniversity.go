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
		SELECT *
		FROM pending_university
		WHERE status = 'approved' AND id = $1
		FOR UPDATE
	`

	err = tx.Get(&pending, fetchQuery, pendingUniversityID)
	if err != nil {
		return 0, richerror.New("manipulationRepository-StabilizeUniversity").WithErr(err)
	}

	switch pending.Action {

	case "create":

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
			VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
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
			return 0, richerror.New("manipulationRepository-StabilizeUniversity").WithErr(err).WithKind(richerror.KindUnexpected).WithMessage("error on insert university")
		}

	case "update":

		updateQuery := `
			UPDATE university
			SET
				name = $1,
				name_english = $2,
				city = $3,
				category = $4,
				image_url = $5,
				description = $6,
				description_english = $7
			WHERE id = $8
		`

		_, err = tx.Exec(
			updateQuery,
			pending.Name,
			pending.NameEnglish,
			pending.City,
			pending.Category,
			pending.ImageUrl,
			pending.Description,
			pending.DescriptionEnglish,
			pending.TargetId,
		)
		if err != nil {
			return 0, richerror.New("manipulationRepository-StabilizeUniversity").WithErr(err).WithKind(richerror.KindUnexpected).WithMessage("error on update university")
		}

	default:
		return 0, richerror.New("manipulationRepository-StabilizeUniversity").WithKind(richerror.KindUnexpected).WithMessage("invalid pending action")
	}

	_, err = tx.Exec(`DELETE FROM pending_university WHERE id = $1`, pendingUniversityID)
	if err != nil {
		return 0, richerror.New("manipulationRepository-StabilizeUniversity").WithErr(err).WithKind(richerror.KindUnexpected).WithMessage("error on delete pending university")
	}

	return pending.SubmittedBy, nil
}
