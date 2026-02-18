package entity

type Professor struct {
	Id                 string            `json:"id"`
	Name               string            `json:"name"`
	NameEnglish        string            `json:"name_english"`
	Description        string            `json:"description"`
	DescriptionEnglish string            `json:"description_english"`
	EducationHistory   map[string]string `json:"education_history"`
	ImageUrl           string            `json:"image_url"`
	RegisteredBy       string            `json:"registered_by"`
}
