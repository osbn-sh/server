package academicService

import (
	"ostadbun/param/academicParam"
)

func (s Service) AddPreRequitesLesson(r academicParam.AddRequiteLesson) error {
	return s.academicRepo.AddPendingLessonPreRequisites(r.LessonID, r.TargetID)
}

func (s Service) AddCoRequitesLesson(r academicParam.AddRequiteLesson) error {
	return s.academicRepo.AddPendingLessonCoRequisites(r.LessonID, r.TargetID)
}
