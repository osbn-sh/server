package manipulationRepository

import (
	"encoding/json"
	"ostadbun/entity"
	"ostadbun/pkg/richerror"
)

// GetProfessorPending returns all professors with 'pending' status
func (d DB) GetProfessorPending() ([]entity.PendingProfessor, error) {
	query := `
        SELECT 
            id,
            name,
            education_history,
            image_url,
            description,
            status,
            submitted_by,
            submitted_at,
            name_english,
            description_english,
            approved_by,
            approved_at,
            rejection_reason
        FROM pending_professor
        WHERE status = 'pending'
    `

	rows, err := d.conn.Conn().Query(query)
	if err != nil {
		return nil, richerror.New("manipulationRepository-GetProfessorPending").WithErr(err).WithKind(richerror.KindUnexpected).WithMessage("error on query pending professor")
	}
	defer rows.Close()

	var professors []entity.PendingProfessor

	for rows.Next() {
		var professor entity.PendingProfessor
		var eduHistoryJSON []byte // موقت برای اسکن داده jsonb

		err := rows.Scan(
			&professor.Id,
			&professor.Name,
			&eduHistoryJSON, // اسکن به []byte
			&professor.ImageUrl,
			&professor.Description,
			&professor.Status,
			&professor.SubmittedBy,
			&professor.SubmittedAt,
			&professor.NameEnglish,
			&professor.DescriptionEnglish,
			&professor.ApprovedBy,
			&professor.ApprovedAt,
			&professor.RejectionReason,
		)
		if err != nil {
			return nil, richerror.New("manipulationRepository-GetProfessorPending").WithErr(err).WithKind(richerror.KindUnexpected).WithMessage("error on query pending professor")
		}

		// تبدیل []byte به *map[string]string
		if eduHistoryJSON != nil {
			var eduMap json.RawMessage
			if err := json.Unmarshal(eduHistoryJSON, &eduMap); err != nil {
				return nil, richerror.New("manipulationRepository-GetProfessorPending").WithErr(err).WithKind(richerror.KindUnexpected).WithMessage("error on unmarshal education history")
			}

			professor.EducationHistory = eduMap
		} // اگر nil باشد، مقدار پیش‌فرض *map[string]string همان nil است

		professors = append(professors, professor)
	}

	if err := rows.Err(); err != nil {
		return nil, richerror.New("manipulationRepository-GetProfessorPending").WithErr(err).WithKind(richerror.KindUnexpected).WithMessage("error on query pending professor")
	}

	return professors, nil
}
