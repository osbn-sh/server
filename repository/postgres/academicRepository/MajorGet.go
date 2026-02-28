package academicRepository

import (
	"ostadbun/entity"
)

func (d DB) MajorGet(id int) (*entity.Major, error) {
	var major entity.Major

	// Query برای جستجوی درس‌ها
	query := `
        SELECT id,name ,name_english,description ,description_english
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
	)

	if errT != nil {
		return nil, errT
	}

	return &major, nil
}
