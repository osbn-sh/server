package academicRepository

import (
	"ostadbun/pkg/richerror"
)

func (d DB) AddPendingLessonPreRequisites(lessonID, TargetID int) error {

	query := `
				insert into lesson_pre_requisite (lesson_id, prerequisite_lesson_id)
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
