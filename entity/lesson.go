package entity

type Lesson struct {
	Id                 string `json:"id"`
	Name               string `json:"name"`
	NameEnglish        string `json:"name_english"`
	Difficulty         int    `json:"difficulty"`
	Description        string `json:"description"`
	DescriptionEnglish string `json:"description_english"`
	RegisteredBy       string `json:"registered_by"`
	IsReleased         bool   `json:"is_released"`
}
