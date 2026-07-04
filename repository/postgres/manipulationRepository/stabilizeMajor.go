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
        SELECT
            name,
            name_english,
            description,
            description_english,
            submitted_by
        FROM pending_major
        WHERE status = 'approved' AND id = $1
        FOR UPDATE
    `

	err = tx.Get(&pending, fetchQuery, pendingMajorID)
	if err != nil {
		return 0, richerror.New("manipulationRepository-StabilizeMajor").WithErr(err).WithKind(richerror.KindUnexpected).WithMessage("error on get pending major")
	}

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
		return 0, richerror.New("manipulationRepository-StabilizeMajor").WithErr(err).WithKind(richerror.KindUnexpected).WithMessage("error on insert pending major")
	}

	deleteQuery := `
		DELETE FROM pending_major
		WHERE id = $1
	`

	_, errE := tx.Exec(deleteQuery, pendingMajorID)
	if errE != nil {
		return 0, richerror.New("manipulationRepository-StabilizeMajor").WithErr(err).WithKind(richerror.KindUnexpected).WithMessage("error on delete pending major")
	}

	return pending.SubmittedBy, nil
}
