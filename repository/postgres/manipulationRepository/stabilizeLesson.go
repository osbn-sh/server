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
		SELECT * FROM pending_lesson
		WHERE status = 'approved' AND id = $1
		FOR UPDATE
	`

	err = tx.Get(&pending, fetchQuery, pendingLessonID)
	if err != nil {
		return 0, richerror.New("manipulationRepository-StabilizeLesson").WithErr(err).WithKind(richerror.KindUnexpected).WithMessage("error on get pending lesson")
	}

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
		VALUES ($1, $2, $3, $4, $5,$6, $7,$8)
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
		return 0, richerror.New("manipulationRepository-StabilizeLesson").WithErr(err).WithKind(richerror.KindUnexpected).WithMessage("error on insert pending lesson")
	}

	_, errE := tx.Exec(`DELETE FROM pending_lesson WHERE id = $1`, pendingLessonID)

	if errE != nil {
		return 0, richerror.New("manipulationRepository-StabilizeLesson").WithErr(err).WithKind(richerror.KindUnexpected).WithMessage("error on delete pending lesson")
	}
	return pending.SubmittedBy, nil
}
