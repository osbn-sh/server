package academicRepository

import (
	"ostadbun/entity"
)

func (d DB) ProfessorGet(id int) (*entity.Professor, error) {
	var professor entity.Professor

	// Query برای جستجوی درس‌ها
	query := `
        SELECT name ,name_english,description ,description_english,education_history,image_url,id
        FROM professor 
        WHERE 
		 	id = $1; 
    `

	// اجرای Query و دریافت نتایج
	errT := d.conn.Conn().QueryRow(query, id).Scan(
		&professor.Name,
		&professor.NameEnglish,
		&professor.Description,
		&professor.DescriptionEnglish,
		&professor.EducationHistory,
		&professor.ImageUrl,
		&professor.Id,
	)

	if errT != nil {
		return nil, errT
	}

	return &professor, nil
}
