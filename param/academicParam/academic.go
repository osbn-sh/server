package academicParam

type RequiteApprovement struct {
	LessonID    int     `json:"lesson_id"`
	TargetID    int     `json:"target_id"`
	CategorySTR string  `json:"category_str"`
	Reason      *string `json:"reason"`
}

type AddRequiteLesson struct {
	LessonID int `json:"lesson_id"`
	TargetID int `json:"target_id"`
}
