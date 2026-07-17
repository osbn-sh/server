package studentRepository

import (
	"fmt"
	"ostadbun/entity"

	"ostadbun/pkg/richerror"
)

func (d DB) GetPass(userID int) ([]entity.PassedLessonInfo, error) {

	query := `
		SELECT p.id, p.name, l.id, l.name, m.id, m.name, u.id, u.name
		FROM passed_lesson_professor_user AS plpu
		JOIN professor  AS p ON plpu.professor_id  = p.id
		JOIN lesson     AS l ON plpu.lesson_id     = l.id
		JOIN major      AS m ON plpu.major_id      = m.id
		JOIN university AS u ON plpu.university_id = u.id
		WHERE plpu.user_id = $1
	`

	rows, err := d.conn.Conn().Query(query, userID)
	if err != nil {
		return nil, richerror.New("studentRepository-GetPass").WithErr(err)
	}
	defer rows.Close()

	var result []entity.PassedLessonInfo
	for rows.Next() {
		var info entity.PassedLessonInfo
		if err := rows.Scan(
			&info.ProfessorID, &info.ProfessorName,
			&info.LessonID, &info.LessonName,
			&info.MajorID, &info.MajorName,
			&info.UniversityID, &info.UniversityName,
		); err != nil {
			return nil, richerror.New("studentRepository-GetPass").WithErr(err)
		}
		result = append(result, info)
	}

	if err := rows.Err(); err != nil {
		return nil, richerror.New("studentRepository-GetPass").WithErr(err)
	}

	if len(result) == 0 {
		return nil, richerror.New("studentRepository-GetPass").
			WithErr(fmt.Errorf("no passed lessons found for user %d", userID))
	}

	return result, nil
}
