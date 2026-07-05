package manipulationRepository

import (
	"ostadbun/entity"
	"ostadbun/pkg/richerror"
)

func (d DB) AddLessonPending(lesson entity.PendingLesson, userId int) error {

	query := `
        insert into pending_lesson(
                                   name,
                                   difficulty,
                                   description,
                                   name_english,
                                   description_english,
                                   term,
                                   action,
                                   target_id,
                                   submitted_by
                                   ) values ($1, $2, $3,$4,$5,$6,$7,$8,$9)
    `

	err := d.conn.Conn().QueryRow(query,
		lesson.Name,
		lesson.Difficulty,
		lesson.Description,
		lesson.NameEnglish,
		lesson.DescriptionEnglish,
		lesson.Term,
		lesson.Action,
		lesson.TargetId,
		userId,
	).Err()

	if err != nil {
		return richerror.New("manipulationRepository-EditLessonPending").WithErr(err)
	}

	return nil
}
