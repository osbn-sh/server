package academicRepository

import (
	"ostadbun/entity"
	"ostadbun/pkg/richerror"
)

func (d DB) ProfessorGet(id int) (*entity.Professor, error) {
	var professor entity.Professor

	// Query برای جستجوی درس‌ها
	query := `
        SELECT id, name ,name_english,description ,description_english,education_history,image_url,href
        FROM professor 
        WHERE 
		 	id = $1; 
    `

	// اجرای Query و دریافت نتایج
	errT := d.conn.Conn().QueryRow(query, id).Scan(
		&professor.Id,
		&professor.Name,
		&professor.NameEnglish,
		&professor.Description,
		&professor.DescriptionEnglish,
		&professor.EducationHistory,
		&professor.ImageUrl,
		&professor.Href,
	)

	if errT != nil {
		return nil, richerror.New("academicRepository-ProfessorGet").WithErr(errT).WithKind(richerror.KindUnexpected)
	}

	return &professor, nil
}

func (d DB) ProfessorGetByHref(href string) (*entity.Professor, error) {
	var professor entity.Professor

	// Query برای جستجوی درس‌ها
	query := `
        SELECT id, name ,name_english,description ,description_english,education_history,image_url,href
        FROM professor 
        WHERE 
		 	href = $1; 
    `

	// اجرای Query و دریافت نتایج
	errT := d.conn.Conn().QueryRow(query, href).Scan(
		&professor.Id,
		&professor.Name,
		&professor.NameEnglish,
		&professor.Description,
		&professor.DescriptionEnglish,
		&professor.EducationHistory,
		&professor.ImageUrl,
		&professor.Href,
	)

	if errT != nil {
		return nil, richerror.New("academicRepository-ProfessorGet").WithErr(errT).WithKind(richerror.KindUnexpected)
	}

	return &professor, nil
}
