package academicRepository

import (
	"fmt"
	"strings" // برای ساختن placeholder ها

	"ostadbun/entity"
	"ostadbun/pkg/richerror"
)

func (d DB) LessonGetMany(ids ...int) (*[]entity.Lesson, error) {
	if len(ids) == 0 {
		return nil, richerror.New("academicRepository-LessonGetMany").WithMessage("no ids provided")
	}

	// ساختن placeholder ها برای شرط IN
	placeholders := make([]string, len(ids))
	for i := range ids {
		placeholders[i] = fmt.Sprintf("$%d", i+1)
	}
	placeholderString := strings.Join(placeholders, ",")

	// کوئری برای جستجوی درس‌ها با شرط IN
	query := fmt.Sprintf(`
        SELECT id, name, name_english, difficulty, description, description_english, term
        FROM lesson
        WHERE id IN (%s);
    `, placeholderString)

	// تبدیل slice of int به slice of interface{} برای ارسال به Query
	args := make([]interface{}, len(ids))
	for i, id := range ids {
		args[i] = id
	}

	rows, err := d.conn.Conn().Query(query, args...) // استفاده از Query به جای QueryRow
	if err != nil {
		return nil, richerror.New("academicRepository-LessonGetMany").WithErr(err).WithKind(richerror.KindUnexpected).WithMessage("error on query multiple rows")
	}
	defer rows.Close()

	var lessons []entity.Lesson // اسلایس برای نگهداری درس‌ها
	for rows.Next() {
		var lesson entity.Lesson
		errScan := rows.Scan(
			&lesson.Id,
			&lesson.Name,
			&lesson.NameEnglish,
			&lesson.Difficulty,
			&lesson.Description,
			&lesson.DescriptionEnglish,
			&lesson.Term,
		)
		if errScan != nil {
			return nil, richerror.New("academicRepository-LessonGetMany").WithErr(errScan).WithKind(richerror.KindUnexpected).WithMessage("error scanning lesson row")
		}
		lessons = append(lessons, lesson)
	}

	// بررسی خطاهای احتمالی در طول تکرار روی rows
	if err = rows.Err(); err != nil {
		return nil, richerror.New("academicRepository-LessonGetMany").WithErr(err).WithKind(richerror.KindUnexpected).WithMessage("error after iterating through rows")
	}

	return &lessons, nil
}
