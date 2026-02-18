package manipulationRepository

import (
	"ostadbun/entity"
)

func (d DB) StabilizeProfessor(pendingProfessorID int) (err error) {

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

	var pending entity.PendingProfessor

	fetchQuery := `
		SELECT * FROM pending_professor
		WHERE status = 'approved' AND id = $1
		FOR UPDATE
	`

	err = tx.Get(&pending, fetchQuery, pendingProfessorID)
	if err != nil {
		return err
	}

	insertQuery := `
		INSERT INTO professor (
			name,
			name_english,
			description,
			description_english,
			education_history,
			image_url,
			registered_by
		)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
	`

	_, err = tx.Exec(
		insertQuery,
		pending.Name,
		pending.NameEnglish,
		pending.Description,
		pending.DescriptionEnglish,
		pending.EducationHistory,
		pending.ImageUrl,
		pending.SubmittedBy,
	)
	if err != nil {
		return err
	}

	_, err = tx.Exec(`DELETE FROM pending_professor WHERE id = $1`, pendingProfessorID)
	return err
}
