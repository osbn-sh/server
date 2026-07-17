package studentService

import (
	"context"
	"ostadbun/entity"
	"ostadbun/param/studentparam"
	Activityconstants "ostadbun/pkg/constants"
)

func (s Service) Remove(id, userID int) (bool, error) {

	return s.academicRepo.RemovePass(id, userID)

}

func (s Service) Add(userID int, student studentparam.StudentPassDetail, c context.Context) error {

	er := s.academicRepo.AddPass(userID, student)

	if er == nil {
		s.activity.Trigger(c, userID, Activityconstants.TriggerAddPassedLesson)
	}

	return er
}

func (s Service) Get(userID int) ([]entity.PassedLessonInfo, error) {

	return s.academicRepo.GetPass(userID)

}
