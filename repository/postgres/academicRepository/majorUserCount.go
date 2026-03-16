package academicRepository

import (
	"fmt"
	"ostadbun/pkg/richerror"
)

func (d DB) UserCountMajor(id int) (int, error) {
	var count int

	// Query برای جستجوی درس‌ها
	query := `
       SELECT count(*) 
        FROM users 
        WHERE 
		 	major_id = $1; 
    `

	// اجرای Query و دریافت نتایج
	errT := d.conn.Conn().QueryRow(query, id).Scan(&count)

	if errT != nil {
		fmt.Println(errT)
		return 0, richerror.New("academicRepository-LessonGet").WithErr(errT).WithKind(richerror.KindUnexpected).WithMessage("error on query row")
	}

	return count, nil
}
