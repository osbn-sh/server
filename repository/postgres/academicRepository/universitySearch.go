package academicRepository

import (
	"ostadbun/entity"
)

func (d DB) UniversitySearch(name string) ([]entity.University, error) {
	var universities []entity.University
	name = "%" + name + "%"
	query := `
			SELECT name, name_english, city, category, image_url,
				   description, description_english
			FROM university
			WHERE
			    name ILIKE $1 OR
				name_english ILIKE $1 OR
				city ILIKE $1 OR
				category ILIKE $1 OR
				description ILIKE $1 OR
				description_english ILIKE $1;
`

	// اجرای query
	rows, err := d.conn.Conn().Query(query, name)
	if err != nil {
		return nil, err
	}

	defer rows.Close() // بستن نتایج پس از پایان

	// پر کردن لیست دانشگاه‌ها
	for rows.Next() {
		var university entity.University
		err := rows.Scan(
			&university.Name,
			&university.NameEnglish,
			&university.City,
			&university.Category,
			&university.ImageUrl,
			&university.Description,
			&university.DescriptionEnglish,
		)
		if err != nil {
			return nil, err // در صورت خطا در Scan، خطا را بازگردانی کن
		}
		universities = append(universities, university)
	}

	// بررسی خطا در حین پیمایش ردیف‌ها
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return universities, nil
}
