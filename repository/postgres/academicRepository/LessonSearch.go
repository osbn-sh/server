package academicRepository

import (
	"ostadbun/entity"
	"ostadbun/pkg/richerror"
)

func (d DB) LessonSearch(name string) ([]entity.Lesson, error) {
	var lessons []entity.Lesson
	name = "%" + name + "%"

	// Query برای جستجوی درس‌ها
	query := `
        SELECT id, name,name_english, difficulty, description ,description_english
        FROM lesson 
        WHERE 
            name ILIKE $1 OR 
			name_english ILIKE $1 OR
			description ILIKE $1 OR
		 	description_english ILIKE $1; 
    `

	// اجرای Query و دریافت نتایج
	rows, err := d.conn.Conn().Query(query, name)
	if err != nil {
		return nil, err // در صورت خطا، خطا را بازگردانی کن
	}
	defer rows.Close() // بستن نتایج پس از پایان

	// پر کردن لیست درس‌ها
	for rows.Next() {
		var lesson entity.Lesson
		err := rows.Scan(
			&lesson.Id,
			&lesson.Name,
			&lesson.NameEnglish,
			&lesson.Difficulty,
			&lesson.Description,
			&lesson.DescriptionEnglish,
		)
		if err != nil {
			return nil, richerror.New("academicRepository-LessonSearch").WithErr(err).WithKind(richerror.KindUnexpected).WithMessage("error on query row") // در صورت خطا در Scan، خطا را بازگردانی کن
		}
		lessons = append(lessons, lesson)
	}

	// بررسی خطا در حین پیمایش ردیف‌ها
	if err := rows.Err(); err != nil {
		return nil, richerror.New("academicRepository-LessonSearch").WithErr(err).WithKind(richerror.KindUnexpected).WithMessage("error on query row")
	}

	return lessons, nil
}
