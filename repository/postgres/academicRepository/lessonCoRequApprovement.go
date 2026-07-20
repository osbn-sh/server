package academicRepository

import (
	"fmt"
	"ostadbun/pkg/richerror"
)

func (d DB) CoRequitesApprovement(id, targetId int, status bool, reason string) error {

	query := `
       	update lesson_co_requisite
       	set 
       	    status = $2,
			rejection_reason = COALESCE($3, rejection_reason)
			where lesson_id = $1 and co_requisite_lesson_id = $4
    `

	res, errT := d.conn.Conn().Exec(query,
		id,
		statusToString(status),
		reason,
		targetId,
	)

	if errT != nil {
		return richerror.New("voteRepository-addoption").WithErr(errT)
	}
	fmt.Println(res.RowsAffected())
	return nil
}

func statusToString(status bool) string {
	if status {
		return "approved"
	}
	return "rejected"
}
