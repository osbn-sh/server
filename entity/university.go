package entity

type University struct {
	Id                 string          `json:"id"`
	Name               string          `json:"name"`
	NameEnglish        *string         `json:"name_english"`
	City               string          `json:"city"`
	Category           string          `json:"category"`
	ImageUrl           string          `json:"image_url"`
	Description        string          `json:"description"`
	DescriptionEnglish *string         `json:"description_english"`
	RegisteredBy       string          `json:"registered_by"`
	Status             string          `json:"status"`
	UsersCount         int             `json:"users_count,omitempty" db:"users_count"`
	Relationships      *MultiDepondMap `json:"relationships,omitempty" db:"relationships"`
}
