package manipulationRepository

import (
	"ostadbun/entity"
	"ostadbun/pkg/richerror"
)

func (d DB) AddLessonPending(lesson entity.PendingLesson, userId int) error {

	query := `
        insert into pending_lesson(
                                   name ,
                                   difficulty,
                                   description,
                                   name_english,
                                   description_english,
                                   term,
                                   submitted_by
                                   ) values ($1, $2, $3,$4,$5,$6,$7)
    `

	err := d.conn.Conn().QueryRow(query,
		lesson.Name,
		lesson.Difficulty,
		lesson.Description,
		lesson.NameEnglish,
		lesson.DescriptionEnglish,
		lesson.Term,
		userId,
	).Err()

	if err != nil {
		return richerror.New("manipulationRepository-AddUniversityPending").WithErr(err)
	}

	return nil
}
