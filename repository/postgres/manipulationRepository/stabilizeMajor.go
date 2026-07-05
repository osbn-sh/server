package manipulationRepository

import (
	"ostadbun/entity"
	"ostadbun/pkg/richerror"
)

func (d DB) StabilizeMajor(pendingMajorID int) (submitterID int64, err error) {

	tx, err := d.conn.Conn().Beginx()
	if err != nil {
		return 0, richerror.New("manipulationRepository-StabilizeMajor").WithErr(err).WithKind(richerror.KindUnexpected).WithMessage("error on begin transaction")
	}

	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			err = tx.Commit()
		}
	}()

	var pending entity.PendingMajor

	fetchQuery := `
        SELECT *
        FROM pending_major
        WHERE status = 'approved' AND id = $1
        FOR UPDATE
    `

	err = tx.Get(&pending, fetchQuery, pendingMajorID)
	if err != nil {
		return 0, richerror.New("manipulationRepository-StabilizeMajor").WithErr(err).WithKind(richerror.KindUnexpected).WithMessage("error on get pending major")
	}

	switch pending.Action {

	case "create":

		insertQuery := `
			INSERT INTO major (
				name,
				name_english,
				description,
				description_english,
				registered_by
			)
			VALUES ($1, $2, $3, $4, $5)
		`

		_, err = tx.Exec(
			insertQuery,
			pending.Name,
			pending.NameEnglish,
			pending.Description,
			pending.DescriptionEnglish,
			pending.SubmittedBy,
		)
		if err != nil {
			return 0, richerror.New("manipulationRepository-StabilizeMajor").WithErr(err).WithKind(richerror.KindUnexpected).WithMessage("error on insert major")
		}

	case "update":

		updateQuery := `
			UPDATE major
			SET
				name = $1,
				name_english = $2,
				description = $3,
				description_english = $4
			WHERE id = $5
		`

		_, err = tx.Exec(
			updateQuery,
			pending.Name,
			pending.NameEnglish,
			pending.Description,
			pending.DescriptionEnglish,
			pending.TargetId,
		)
		if err != nil {
			return 0, richerror.New("manipulationRepository-StabilizeMajor").WithErr(err).WithKind(richerror.KindUnexpected).WithMessage("error on update major")
		}

	default:
		return 0, richerror.New("manipulationRepository-StabilizeMajor").WithKind(richerror.KindUnexpected).WithMessage("invalid pending action")
	}

	_, err = tx.Exec(`DELETE FROM pending_major WHERE id = $1`, pendingMajorID)
	if err != nil {
		return 0, richerror.New("manipulationRepository-StabilizeMajor").WithErr(err).WithKind(richerror.KindUnexpected).WithMessage("error on delete pending major")
	}

	return pending.SubmittedBy, nil
}
