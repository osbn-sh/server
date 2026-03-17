package entity

type Lesson struct {
	Id                 string          `json:"id"`
	Name               string          `json:"name"`
	NameEnglish        string          `json:"name_english"`
	Term               string          `json:"term"`
	Difficulty         int             `json:"difficulty"`
	Description        string          `json:"description"`
	DescriptionEnglish string          `json:"description_english"`
	RegisteredBy       string          `json:"registered_by"`
	IsReleased         bool            `json:"is_released"`
	UsersCount         int             `json:"users_count" db:"users_count"`
	Relationships      *MultiDepondMap `json:"relationships" db:"relationships"`
	PreRequites        *[]Lesson       `json:"pre_requites" db:"pre_requites"`
	CoRequites         *[]Lesson       `json:"co_requites" db:"co_requites"`
}
