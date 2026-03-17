package academicRepository

import (
	"fmt"
	"ostadbun/entity"
	"ostadbun/pkg/richerror"
)

func (d DB) MultiDepending(id int, entityTarget string) (*[]entity.MultiDepending, error) {
	var multi []entity.MultiDepending

	// Query برای جستجوی درس‌ها
	query := `
        SELECT *
        FROM multi_depending 
        WHERE 
		%s = $1; 
    `

	// اجرای Query و دریافت نتایج
	rows, errT := d.conn.Conn().Query(fmt.Sprintf(query, entityTarget), id)

	if errT != nil {
		return nil, richerror.New("academicRepository-MultiDepending").WithErr(errT).WithKind(richerror.KindUnexpected) // در صورت خطا، خطا را بازگردانی کن
	}
	defer rows.Close() // بستن نتایج پس از پایان

	for rows.Next() {
		var multiThis entity.MultiDepending
		err := rows.Scan(
			&multiThis.ProfessorId,
			&multiThis.LessonId,
			&multiThis.UniversityId,
			&multiThis.MajorId,
		)
		if err != nil {
			return nil, richerror.New("academicRepository-majorSearch").WithErr(err).WithKind(richerror.KindUnexpected) // در صورت خطا در Scan، خطا را بازگردانی کن
		}
		multi = append(multi, multiThis)
	}

	// بررسی خطا در حین پیمایش ردیف‌ها
	if err := rows.Err(); err != nil {
		return nil, richerror.New("academicRepository-majorSearch").WithErr(err).WithKind(richerror.KindUnexpected)
	}

	fmt.Println(multi)
	return &multi, nil
}
