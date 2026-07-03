package entity

import (
	"encoding/json"
)

type Professor struct {
	Id                 int               `json:"id"`
	Name               string            `json:"name"`
	NameEnglish        *string           `json:"name_english"`
	Description        *string           `json:"description"`
	DescriptionEnglish *string           `json:"description_english"`
	EducationHistory   json.RawMessage   `json:"education_history"`
	ImageUrl           string            `json:"image_url"`
	RegisteredBy       string            `json:"registered_by"`
	UsersCount         int               `json:"users_count,omitempty" db:"users_count"`
	Href               *string           `json:"href"`
	Relationships      *MultiDependSlice `json:"relationships,omitempty" db:"relationships"`
}
