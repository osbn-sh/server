package academicRepository

import (
	"ostadbun/entity"
	"ostadbun/pkg/richerror"
)

func (d DB) LessonGet(id int) (*entity.Lesson, error) {
	var lessons entity.Lesson

	// Query برای جستجوی درس‌ها
	query := `
        SELECT id,name ,name_english,difficulty,description ,description_english,term,href
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
		&lessons.Href,
	)

	if errT != nil {
		return nil, richerror.New("academicRepository-LessonGet").WithErr(errT).WithKind(richerror.KindUnexpected).WithMessage("error on query row")
	}

	return &lessons, nil
}

func (d DB) LessonGetByHref(href string) (*entity.Lesson, error) {
	var lessons entity.Lesson

	// Query برای جستجوی درس‌ها
	query := `
        SELECT id,name ,name_english,difficulty,description ,description_english,term,href
        FROM lesson 
        WHERE 
		 	href = $1; 
    `

	// اجرای Query و دریافت نتایج
	errT := d.conn.Conn().QueryRow(query, href).Scan(
		&lessons.Id,
		&lessons.Name,
		&lessons.NameEnglish,
		&lessons.Difficulty,
		&lessons.Description,
		&lessons.DescriptionEnglish,
		&lessons.Term,
		&lessons.Href,
	)

	if errT != nil {
		return nil, richerror.New("academicRepository-LessonGet").WithErr(errT).WithKind(richerror.KindUnexpected).WithMessage("error on query row")
	}

	return &lessons, nil
}
