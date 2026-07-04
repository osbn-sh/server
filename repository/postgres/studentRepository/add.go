package studentRepository

import (
	"fmt"
	"ostadbun/param/studentparam"
	"ostadbun/pkg/richerror"
)

func (d DB) AddPass(userID int, data studentparam.StudentPassDetail) error {

	query := `
       	INSERT INTO passed_lesson_professor_user (user_id, university_id, major_id, professor_id, lesson_id) 
		VALUES ($1, $2, $3, $4, $5)
    `

	res, errT := d.conn.Conn().Exec(query,
		userID,
		data.UniversityID,
		data.MajorID,
		data.ProfessorID,
		data.LessonID,
	)

	if errT != nil {
		return richerror.New("academicRepository-UpsertPass").WithErr(errT)
	}
	fmt.Println(res.RowsAffected())
	return nil
}
