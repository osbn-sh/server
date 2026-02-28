package academicRepository

import (
	"ostadbun/entity"
)

func (d DB) UniversityGet(id int) (*entity.University, error) {
	var university entity.University

	// Query برای جستجوی درس‌ها
	query := `
        SELECT id,name ,name_english,description ,description_english,category,image_url
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
	)

	if errT != nil {
		return nil, errT
	}

	return &university, nil
}
