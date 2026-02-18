package manipulationRepository

import "ostadbun/entity"

func (d DB) StabilizeMajor(pendingMajorID int) (err error) {

	tx, err := d.conn.Conn().Beginx()
	if err != nil {
		return err
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
		return err
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
		return err
	}

	deleteQuery := `
		DELETE FROM pending_major
		WHERE id = $1
	`

	_, err = tx.Exec(deleteQuery, pendingMajorID)
	if err != nil {
		return err
	}

	return nil
}
