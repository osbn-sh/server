package academicRepository

import (
	"ostadbun/entity"
	"ostadbun/pkg/richerror"
)

func (d DB) MajorGet(id int) (*entity.Major, error) {
	var major entity.Major

	// Query برای جستجوی درس‌ها
	query := `
        SELECT id,name ,name_english,description ,description_english,href
        FROM major 
        WHERE 
		 	id = $1; 
    `

	// اجرای Query و دریافت نتایج
	errT := d.conn.Conn().QueryRow(query, id).Scan(
		&major.Id,
		&major.Name,
		&major.NameEnglish,
		&major.Description,
		&major.DescriptionEnglish,
		&major.Href,
	)

	if errT != nil {
		return nil, richerror.New("academicRepository-MajorSearch").WithErr(errT).WithKind(richerror.KindUnexpected)
	}

	return &major, nil
}

func (d DB) MajorGetByHref(href string) (*entity.Major, error) {
	var major entity.Major

	// Query برای جستجوی درس‌ها
	query := `
        SELECT id,name ,name_english,description ,description_english,href
        FROM major 
        WHERE 
		 	href = $1; 
    `

	// اجرای Query و دریافت نتایج
	errT := d.conn.Conn().QueryRow(query, href).Scan(
		&major.Id,
		&major.Name,
		&major.NameEnglish,
		&major.Description,
		&major.DescriptionEnglish,
		&major.Href,
	)

	if errT != nil {
		return nil, richerror.New("academicRepository-MajorSearch").WithErr(errT).WithKind(richerror.KindUnexpected)
	}

	return &major, nil
}
