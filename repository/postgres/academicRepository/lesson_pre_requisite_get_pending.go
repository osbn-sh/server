package academicRepository

import (
	"ostadbun/entity"
	"ostadbun/pkg/richerror"
)

func (d DB) GetPendingLessonPreRequisites() ([]entity.PreRequites, error) {

	query := `
SELECT
    lcr.lesson_id,
    lcr.prerequisite_lesson_id,
    l1.name,
    l2.name
FROM lesson_pre_requisite lcr
JOIN lesson l1
    ON l1.id = lcr.lesson_id
JOIN lesson l2
    ON l2.id = lcr.prerequisite_lesson_id
WHERE lcr.status = 'pending'
`

	rows, err := d.conn.Conn().Query(query)
	if err != nil {
		return nil, richerror.New("academicRepository-GetPendingLessonPreRequisites").
			WithErr(err).
			WithKind(richerror.KindUnexpected)
	}
	defer rows.Close()

	var items []entity.PreRequites

	for rows.Next() {
		var item entity.PreRequites

		if err := rows.Scan(
			&item.LessonID,
			&item.PreRequisiteLessonId,
			&item.LessonName,
			&item.PreRequisiteLessonName,
		); err != nil {
			return nil, richerror.New("academicRepository-GetPendingLessonCoRequisites").
				WithErr(err).
				WithKind(richerror.KindUnexpected)
		}

		items = append(items, item)
	}

	if err := rows.Err(); err != nil {
		return nil, richerror.New("academicRepository-GetPendingLessonCoRequisites").
			WithErr(err).
			WithKind(richerror.KindUnexpected)
	}

	return items, nil
}
