package academicRepository

import (
	"ostadbun/entity"
	"ostadbun/pkg/richerror"
)

func (d DB) LessonGet(id int) (*entity.Lesson, error) {
	var lessons entity.Lesson

	// Query برای جستجوی درس‌ها
	query := `
        SELECT id,name ,name_english,difficulty,description ,description_english,term
        FROM lesson 
        WHERE 
		 	id = $1; 
    `

	// اجرای Query و دریافت نتایج
	errT := d.conn.Conn().QueryRow(query, id).Scan(
		&lessons.Id,
		&lessons.Name,
		&lessons.NameEnglish,
		&lessons.Difficulty,
		&lessons.Description,
		&lessons.DescriptionEnglish,
		&lessons.Term,
	)

	if errT != nil {
		return nil, richerror.New("academicRepository-LessonGet").WithErr(errT).WithKind(richerror.KindUnexpected).WithMessage("error on query row")
	}

	return &lessons, nil
}
