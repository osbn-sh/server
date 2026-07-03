package academicRepository

import (
	"fmt"
	"ostadbun/pkg/richerror"
	uniquer "ostadbun/pkg/uniqer"
)

func (d DB) UserCountProfessor(id int) ([]int, error) {

	query := `
       SELECT user_id,professor_id 
        FROM passed_lesson_professor_user 
        WHERE 
		 	professor_id = $1; 
    `

	// اجرای Query و دریافت نتایج
	rows, errT := d.conn.Conn().Query(query, id)

	if errT != nil {

	}

	defer rows.Close()

	var userIDs []int
	for rows.Next() {
		var uid int
		var pid int
		err := rows.Scan(&uid, &pid)
		if err != nil {

		}
		userIDs = append(userIDs, uid)
		fmt.Println(uid, pid)
	}

	if errT != nil {
		fmt.Println(errT)
		return []int{}, richerror.New("academicRepository-UserCountProfessor").WithErr(errT).WithKind(richerror.KindUnexpected).WithMessage("error on query row")
	}

	unique := uniquer.Unique(userIDs)

	return unique, nil
}
