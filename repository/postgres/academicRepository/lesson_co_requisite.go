package academicRepository

import (
	"fmt"
	"ostadbun/pkg/richerror"
)

func (d DB) LessonCoRequisites(id int) ([]int, error) {

	query := `
       SELECT co_requisite_lesson_id 
        FROM lesson_co_requisite 
        WHERE 
		 	lesson_id = $1; 
    `

	// اجرای Query و دریافت نتایج
	rows, errT := d.conn.Conn().Query(query, id)

	if errT != nil {

	}

	defer rows.Close()

	var id_s []int
	for rows.Next() {
		var thisID int
		err := rows.Scan(&thisID)
		if err != nil {

		}
		id_s = append(id_s, thisID)
	}

	if errT != nil {
		fmt.Println(errT)
		return []int{}, richerror.New("academicRepository-UserCountProfessor").WithErr(errT).WithKind(richerror.KindUnexpected).WithMessage("error on query row")
	}

	return id_s, nil
}
