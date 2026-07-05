package manipulationRepository

import (
	"ostadbun/entity"
	"ostadbun/pkg/richerror"
)

func (d DB) AddMajorPending(major entity.PendingMajor, userId int) error {
	query := `
        INSERT INTO pending_major (
            name, 
            description,
            name_english,
            description_english,
			action,
            target_id,
            submitted_by
        ) 
        VALUES ($1, $2, $3, $4, $5,$6,$7)
    `

	err := d.conn.Conn().QueryRow(
		query,
		major.Name,
		major.Description,
		major.NameEnglish,
		major.DescriptionEnglish,
		major.Action,
		major.TargetId,
		userId,
	).Err()

	if err != nil {
		return richerror.New("manipulationRepository-AddUniversityPending").WithErr(err)
	}

	return nil
}
