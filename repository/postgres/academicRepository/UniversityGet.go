package academicRepository

import (
	"ostadbun/entity"
	"ostadbun/pkg/richerror"
)

func (d DB) UniversityGet(id int) (*entity.University, error) {
	var university entity.University

	// Query برای جستجوی درس‌ها
	query := `
        SELECT id,name ,name_english,description ,description_english,category,image_url,href,city
        FROM university 
        WHERE 
		 	id = $1; 
    `

	// اجرای Query و دریافت نتایج
	errT := d.conn.Conn().QueryRow(query, id).Scan(

		&university.Id,
		&university.Name,
		&university.NameEnglish,
		&university.Description,
		&university.DescriptionEnglish,
		&university.Category,
		&university.ImageUrl,
		&university.Href,
		&university.City,
	)

	if errT != nil {
		return nil, richerror.New("academicRepository-UniversityGet").WithErr(errT).WithKind(richerror.KindUnexpected)
	}

	return &university, nil
}

func (d DB) UniversityGetByHref(href string) (*entity.University, error) {
	var university entity.University

	// Query برای جستجوی درس‌ها
	query := `
        SELECT id,name ,name_english,description ,description_english,category,image_url,href
        FROM university 
        WHERE 
		 	href = $1; 
    `

	// اجرای Query و دریافت نتایج
	errT := d.conn.Conn().QueryRow(query, href).Scan(

		&university.Id,
		&university.Name,
		&university.NameEnglish,
		&university.Description,
		&university.DescriptionEnglish,
		&university.Category,
		&university.ImageUrl,
		&university.Href,
	)

	if errT != nil {
		return nil, richerror.New("academicRepository-UniversityGet").WithErr(errT).WithKind(richerror.KindUnexpected)
	}

	return &university, nil
}
