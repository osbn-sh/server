package manipulationService

import (
	"ostadbun/entity"
)

// GetPendingUniversity -1 all
func (m Manipulation) GetPendingUniversity(userId int) ([]entity.PendingUniversity, error) {

	return m.manipulationRepo.GetUniversityPending(userId)

}

// -1 all
func (m Manipulation) GetPendingProfessor(userId int) ([]entity.PendingProfessor, error) {

	return m.manipulationRepo.GetProfessorPending(userId)

}

// -1 all
func (m Manipulation) GetPendingLesson(userId int) ([]entity.PendingLesson, error) {

	return m.manipulationRepo.GetLessonPending(userId)

}

// -1 all
func (m Manipulation) GetPendingMajor(userId int) ([]entity.PendingMajor, error) {

	return m.manipulationRepo.GetMajorPending(userId)

}
