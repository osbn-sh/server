package manipulationParam

import "encoding/json"

type PendingProfessor struct {
	Name               string          `json:"name"`
	EducationHistory   json.RawMessage `json:"education_history"`
	ImageUrl           string          `json:"image_url"`
	Description        string          `json:"description"`
	NameEnglish        string          `json:"name_english"`
	DescriptionEnglish string          `json:"description_english"`
}
