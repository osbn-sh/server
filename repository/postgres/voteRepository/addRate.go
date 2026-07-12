package voteRepository

import (
	"database/sql"
	"errors"
	"ostadbun/param/voteparam"
	"ostadbun/pkg/richerror"
	"time"
)

func (d DB) AddRate(userID int, data voteparam.Vote) error {

	// چک می‌کنیم owner همون option با target ارسالی یکی باشه
	var owner string
	errQ := d.conn.Conn().QueryRow(
		`SELECT owner FROM option WHERE id = $1`,
		data.OptionID,
	).Scan(&owner)

	if errQ != nil {
		if errors.Is(errQ, sql.ErrNoRows) {
			return richerror.New("voteRepository-AddRate").
				WithMessage("option not found").
				WithKind(richerror.KindNotFound)
		}
		return richerror.New("voteRepository-AddRate").WithErr(errQ).WithKind(richerror.KindUnexpected)
	}

	if owner != data.Target {
		return richerror.New("voteRepository-AddRate").
			WithMessage("option owner does not match target").
			WithKind(richerror.KindInvalid)
	}

	universityID, majorID, professorID, lessonID := resolveTarget(data.Target, data.TargetID)

	query := `
		INSERT INTO vote (user_id, option_id, rate, rate_time, university_id, major_id, professor_id, lesson_id)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
	`

	_, errT := d.conn.Conn().Exec(query,
		userID,
		data.OptionID,
		data.Rate,
		time.Now(),
		universityID,
		majorID,
		professorID,
		lessonID,
	)

	if errT != nil {
		return richerror.New("voteRepository-AddRate").WithErr(errT).WithKind(richerror.KindUnexpected)
	}

	return nil
}

// resolveTarget بر اساس نوع target، فقط ستون مربوطه رو مقداردهی می‌کنه و بقیه NULL می‌مونن
func resolveTarget(target string, targetID int) (universityID, majorID, professorID, lessonID sql.NullInt64) {
	id := sql.NullInt64{Int64: int64(targetID), Valid: true}

	switch target {
	case "university":
		universityID = id
	case "major":
		majorID = id
	case "professor":
		professorID = id
	case "lesson":
		lessonID = id
	}

	return
}
