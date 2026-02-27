package entity

import "time"

type Major struct {
	Id                 string     `json:"id" db:"id"`
	Name               string     `json:"name" db:"name"`
	Registered_by      string     `db:"registered_by" json:"registered_by"`
	NameEnglish        *string    `json:"name_english" db:"name_english"`
	SubmittedAt        *time.Time `json:"submitted_at,omitempty" db:"submitted_at"`
	SubmittedBy        string     `json:"submitted_by" db:"submitted_by"`
	Description        *string    `json:"description" db:"description"`
	DescriptionEnglish *string    `json:"description_english" db:"description_english"`
}
