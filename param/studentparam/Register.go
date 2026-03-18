package studentparam

type PassPut struct {
	LessonID    int `json:"lesson_id"`
	ProfessorID int `json:"professor_id"`
}
type StudentPassDetail struct {
	LessonID     int `json:"lesson_id"`
	ProfessorID  int `json:"professor_id"`
	MajorID      int `json:"major_id"`
	UniversityID int `json:"university_id"`
}
