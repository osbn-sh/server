package manipulationRepository

import (
	"ostadbun/entity"
	"ostadbun/pkg/richerror"
)

func (d DB) StabilizeProfessor(pendingProfessorID int) (err error) {

	tx, err := d.conn.Conn().Beginx()
	if err != nil {
		return richerror.New("manipulationRepository-StabilizeProfessor").WithErr(err).WithKind(richerror.KindUnexpected).WithMessage("error on  on begin transaction")
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
		return richerror.New("manipulationRepository-StabilizeProfessor").WithErr(err).WithKind(richerror.KindUnexpected).WithMessage("error on get pending Professor")
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
		return richerror.New("manipulationRepository-StabilizeProfessor").WithErr(err).WithKind(richerror.KindUnexpected).WithMessage("error on insert pending Professor")
	}

	_, errE := tx.Exec(`DELETE FROM pending_professor WHERE id = $1`, pendingProfessorID)

	if errE != nil {
		return richerror.New("manipulationRepository-StabilizeProfessor").WithErr(err).WithKind(richerror.KindUnexpected).WithMessage("error on delete pending Professor")
	}

	return nil
}
