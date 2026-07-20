package academicRepository

import (
	"ostadbun/pkg/richerror"
)

func (d DB) AddPendingLessonCoRequisites(lessonID, TargetID int) error {

	query := `
				insert into lesson_co_requisite (lesson_id, co_requisite_lesson_id)
				values ($1, $2);
`

	err := d.conn.Conn().QueryRow(query,
		lessonID, TargetID,
	).Err()

	if err != nil {
		return richerror.New("manipulationRepository-EditLessonPending").WithErr(err)
	}

	return nil
}
