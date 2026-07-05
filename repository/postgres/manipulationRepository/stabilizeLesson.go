package manipulationRepository

import (
	"ostadbun/entity"
	"ostadbun/pkg/richerror"
)

func (d DB) StabilizeLesson(pendingLessonID int) (submitterID int64, err error) {

	tx, err := d.conn.Conn().Beginx()
	if err != nil {
		return 0, richerror.New("manipulationRepository-StabilizeLesson").WithErr(err).WithKind(richerror.KindUnexpected).WithMessage("error on begin transaction")
	}

	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			err = tx.Commit()
		}
	}()

	var pending entity.PendingLesson

	fetchQuery := `
		SELECT *
		FROM pending_lesson
		WHERE status = 'approved' AND id = $1
		FOR UPDATE
	`

	err = tx.Get(&pending, fetchQuery, pendingLessonID)
	if err != nil {
		return 0, richerror.New("manipulationRepository-StabilizeLesson").WithErr(err).WithKind(richerror.KindUnexpected).WithMessage("error on get pending lesson")
	}

	switch pending.Action {

	case "create":

		insertQuery := `
			INSERT INTO lesson (
				name,
				name_english,
				difficulty,
				description,
				description_english,
				term,
				registered_by,
				is_released
			)
			VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
		`

		_, err = tx.Exec(
			insertQuery,
			pending.Name,
			pending.NameEnglish,
			pending.Difficulty,
			pending.Description,
			pending.DescriptionEnglish,
			pending.Term,
			pending.SubmittedBy,
			true,
		)
		if err != nil {
			return 0, richerror.New("manipulationRepository-StabilizeLesson").WithErr(err).WithKind(richerror.KindUnexpected).WithMessage("error on insert lesson")
		}

	case "update":

		updateQuery := `
			UPDATE lesson
			SET
				name = $1,
				name_english = $2,
				difficulty = $3,
				description = $4,
				description_english = $5,
				term = $6
			WHERE id = $7
		`

		_, err = tx.Exec(
			updateQuery,
			pending.Name,
			pending.NameEnglish,
			pending.Difficulty,
			pending.Description,
			pending.DescriptionEnglish,
			pending.Term,
			pending.TargetId,
		)
		if err != nil {
			return 0, richerror.New("manipulationRepository-StabilizeLesson").WithErr(err).WithKind(richerror.KindUnexpected).WithMessage("error on update lesson")
		}

	default:
		return 0, richerror.New("manipulationRepository-StabilizeLesson").WithKind(richerror.KindUnexpected).WithMessage("invalid pending action")
	}

	_, err = tx.Exec(`DELETE FROM pending_lesson WHERE id = $1`, pendingLessonID)
	if err != nil {
		return 0, richerror.New("manipulationRepository-StabilizeLesson").WithErr(err).WithKind(richerror.KindUnexpected).WithMessage("error on delete pending lesson")
	}

	return pending.SubmittedBy, nil
}
