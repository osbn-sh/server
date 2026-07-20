package entity

type CoRequites struct {
	LessonID              int     `json:"lessonId"`
	CoRequisiteLessonId   int     `json:"co_requisite_lesson_id"`
	LessonName            string  `json:"lessonName"`
	CoRequisiteLessonName string  `json:"co_requisite_lesson_name"`
	RejectionReason       *string `json:"rejectionReason"`
	Status                string  `json:"status"`
}

type PreRequites struct {
	LessonID               int     `json:"lessonId"`
	PreRequisiteLessonId   int     `json:"prerequisite_lesson_id"`
	LessonName             string  `json:"lessonName"`
	PreRequisiteLessonName string  `json:"prerequisite_lesson_name"`
	RejectionReason        *string `json:"rejectionReason"`
	Status                 string  `json:"status"`
}

type OutParam struct {
	PreRequites []PreRequites `json:"pre_requisites,omitempty"`
	CoRequites  []CoRequites  `json:"co_requisites,omitempty"`
}
