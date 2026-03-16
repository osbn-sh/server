package entity

type MultiDepending struct {
	ProfessorId  int `json:"professor_id"`
	LessonId     int `json:"lesson_id"`
	UniversityId int `json:"university_id"`
	MajorId      int `json:"major_id"`
}
type MultiDepondMap struct {
	University map[int]University
	Lessons    map[int]Lesson
	Professor  map[int]Professor
	Major      map[int]Major
}
