package entity

type Lesson struct {
	Id                 int               `json:"id"`
	Name               string            `json:"name"`
	NameEnglish        string            `json:"name_english"`
	Term               string            `json:"term"`
	Difficulty         int               `json:"difficulty"`
	Description        string            `json:"description"`
	DescriptionEnglish string            `json:"description_english"`
	RegisteredBy       string            `json:"registered_by"`
	IsReleased         bool              `json:"is_released"`
	UsersCount         int               `json:"users_count,omitempty" db:"users_count"`
	Relationships      *MultiDependSlice `json:"relationships,omitempty" db:"relationships"`
	PreRequites        *[]Lesson         `json:"pre_requites,omitempty" db:"pre_requites"`
	CoRequites         *[]Lesson         `json:"co_requites,omitempty" db:"co_requites"`
}
