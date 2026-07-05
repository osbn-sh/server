package manipulationRepository

import (
	"ostadbun/entity"
	"ostadbun/pkg/richerror"
)

func (d DB) StabilizeProfessor(pendingProfessorID int) (submitterID int64, err error) {

	tx, err := d.conn.Conn().Beginx()
	if err != nil {
		return 0, richerror.New("manipulationRepository-StabilizeProfessor").WithErr(err).WithKind(richerror.KindUnexpected).WithMessage("error on begin transaction")
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
		SELECT *
		FROM pending_professor
		WHERE status = 'approved' AND id = $1
		FOR UPDATE
	`

	err = tx.Get(&pending, fetchQuery, pendingProfessorID)
	if err != nil {
		return 0, richerror.New("manipulationRepository-StabilizeProfessor").WithErr(err).WithKind(richerror.KindUnexpected).WithMessage("error on get pending professor")
	}

	switch pending.Action {

	case "create":

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
			return 0, richerror.New("manipulationRepository-StabilizeProfessor").WithErr(err).WithKind(richerror.KindUnexpected).WithMessage("error on insert professor")
		}

	case "update":

		updateQuery := `
			UPDATE professor
			SET
				name = $1,
				name_english = $2,
				description = $3,
				description_english = $4,
				education_history = $5,
				image_url = $6
			WHERE id = $7
		`

		_, err = tx.Exec(
			updateQuery,
			pending.Name,
			pending.NameEnglish,
			pending.Description,
			pending.DescriptionEnglish,
			pending.EducationHistory,
			pending.ImageUrl,
			pending.TargetId,
		)
		if err != nil {
			return 0, richerror.New("manipulationRepository-StabilizeProfessor").WithErr(err).WithKind(richerror.KindUnexpected).WithMessage("error on update professor")
		}

	default:
		return 0, richerror.New("manipulationRepository-StabilizeProfessor").WithKind(richerror.KindUnexpected).WithMessage("invalid pending action")
	}

	_, err = tx.Exec(`DELETE FROM pending_professor WHERE id = $1`, pendingProfessorID)
	if err != nil {
		return 0, richerror.New("manipulationRepository-StabilizeProfessor").WithErr(err).WithKind(richerror.KindUnexpected).WithMessage("error on delete pending professor")
	}

	return pending.SubmittedBy, nil
}
