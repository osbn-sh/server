package manipulationRepository

import (
	"database/sql"
	"ostadbun/entity"
	"ostadbun/pkg/richerror"
)

// GetLessonPending returns all lessons with 'pending' status
func (d DB) GetLessonPending(userId int) ([]entity.PendingLesson, error) {
	query := `
        SELECT 
            id,
            name,
            name_english,
            description_english,
            difficulty,
            description,
            term,
            status,
            submitted_by,
            submitted_at,
            approved_by,
            approved_at,
            rejection_reason
        FROM pending_lesson
        WHERE status = 'pending'
    `

	var (
		rows *sql.Rows
		err  error
	)

	if userId > 0 {
		query += " AND submitted_by = $1"
		rows, err = d.conn.Conn().Query(query, userId)
	} else {
		rows, err = d.conn.Conn().Query(query)
	}
	if err != nil {
		return nil, richerror.New("manipulationRepository-GetLessonPending").WithErr(err).WithKind(richerror.KindUnexpected).WithMessage("error on query pending lesson")
	}
	defer rows.Close()

	var lessons []entity.PendingLesson

	for rows.Next() {
		var lesson entity.PendingLesson
		err := rows.Scan(
			&lesson.Id,
			&lesson.Name,
			&lesson.NameEnglish,
			&lesson.DescriptionEnglish,
			&lesson.Difficulty,
			&lesson.Description,
			&lesson.Term,
			&lesson.Status,
			&lesson.SubmittedBy,
			&lesson.SubmittedAt,
			&lesson.ApprovedBy,
			&lesson.ApprovedAt,
			&lesson.RejectionReason,
		)
		if err != nil {
			return nil, richerror.New("manipulationRepository-GetLessonPending").WithErr(err).WithKind(richerror.KindUnexpected).WithMessage("error on query pending lesson")
		}
		lessons = append(lessons, lesson)
	}

	if err := rows.Err(); err != nil {
		return nil, richerror.New("manipulationRepository-GetLessonPending").WithErr(err).WithKind(richerror.KindUnexpected).WithMessage("error on rows scan")
	}

	return lessons, nil
}
