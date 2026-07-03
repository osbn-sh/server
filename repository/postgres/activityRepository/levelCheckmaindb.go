package activityRepository

import "ostadbun/pkg/richerror"

func (d DB) MainStoreCalculateAndFetch(userid int) (int, error) {
	var totalValue int

	query := `
        SELECT COALESCE(SUM(a.value), 0)
        FROM activity_history ah
        JOIN activity a ON ah.activity_id = a.id
        WHERE ah.user_id = $1
    `

	err := d.conn.Conn().QueryRow(query, userid).Scan(&totalValue)

	if err != nil {
		return -1, richerror.New("activityRepository-MainStoreCalculateAndFetch").WithErr(err).WithKind(richerror.KindUnexpected).WithMessage("error on query row")
	}

	return totalValue, nil
}
