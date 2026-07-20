package academicService

import (
	"ostadbun/param/academicParam"
)

func (s Service) ApprovementPreRequites(r academicParam.RequiteApprovement, status bool) error {
	return s.academicRepo.PreRequitesApprovement(r.LessonID, r.TargetID, status, r.Reason)
}

func (s Service) ApprovementCoRequites(r academicParam.RequiteApprovement, status bool) error {
	return s.academicRepo.CoRequitesApprovement(r.LessonID, r.TargetID, status, r.Reason)
}
