package academicRepository

import (
	"ostadbun/entity"
	"ostadbun/pkg/richerror"
)

func (d DB) GetPendingLessonCoRequisites() ([]entity.CoRequites, error) {

	query := `
SELECT
    lcr.lesson_id,
    lcr.co_requisite_lesson_id,
    l1.name,
    l2.name
FROM lesson_co_requisite lcr
JOIN lesson l1
    ON l1.id = lcr.lesson_id
JOIN lesson l2
    ON l2.id = lcr.co_requisite_lesson_id
WHERE lcr.status = 'pending'
`

	rows, err := d.conn.Conn().Query(query)
	if err != nil {
		return nil, richerror.New("academicRepository-GetPendingLessonCoRequisites").
			WithErr(err).
			WithKind(richerror.KindUnexpected)
	}
	defer rows.Close()

	var items []entity.CoRequites

	for rows.Next() {
		var item entity.CoRequites

		if err := rows.Scan(
			&item.LessonID,
			&item.CoRequisiteLessonId,
			&item.LessonName,
			&item.CoRequisiteLessonName,
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
