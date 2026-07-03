package studentRepository

import (
	"fmt"
	"ostadbun/pkg/richerror"
)

func (d DB) RemovePass(id, userID int) (bool, error) {

	fmt.Println("312dw", userID, id)
	query := `
       	DELETE FROM passed_lesson_professor_user WHERE id = $1 AND user_id = $2;
    `
	res, errT := d.conn.Conn().Exec(query,
		id, userID,
	)

	if errT != nil {
		fmt.Printf("Error during removePass operation: %v\n", errT)
		return false, richerror.New("academicRepository-RemovePass").WithErr(errT).WithKind(richerror.KindUnexpected).WithMessage("error during RemovePass operation")
	}

	n, errR := res.RowsAffected()
	if errR != nil {
		fmt.Printf("Error during RemovePass operation: %v\n", errR)
	}
	return n > 0, nil
}
