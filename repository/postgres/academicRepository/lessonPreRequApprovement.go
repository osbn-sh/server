package academicRepository

import (
	"fmt"
	"ostadbun/pkg/richerror"
)

func (d DB) PreRequitesApprovement(id, targetId int, status bool, reason *string) error {

	query := `
       	update lesson_pre_requisite
       	set 
       	    status = $2,
       		rejection_reason = COALESCE($3, rejection_reason)
			where lesson_id = $1 and prerequisite_lesson_id = $4
    `

	res, errT := d.conn.Conn().Exec(query,
		id,
		statusToString(status),
		reason,
		targetId,
	)

	fmt.Println(errT)
	if errT != nil {
		return richerror.New("voteRepository-addoption").WithErr(errT)
	}
	fmt.Println(res.RowsAffected())
	return nil
}
