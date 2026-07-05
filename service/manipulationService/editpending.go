package manipulationService

import (
	"fmt"
	"ostadbun/entity"
	"ostadbun/pkg/richerror"
)

func (m Manipulation) EditPendingUniversity(lesson entity.PendingUniversity, userId int) error {

	return m.manipulationRepo.AddUniversityPending(lesson, userId)

}

func (m Manipulation) EditPendingProfessor(lesson entity.PendingProfessor, userId int) error {

	return m.manipulationRepo.AddProfessorPending(lesson, userId)

}

func (m Manipulation) EditPendingLesson(lesson entity.PendingLesson, userId int) error {

	if lesson.Difficulty > 5 || lesson.Difficulty < 1 {
		return richerror.New("addpendinglesson.service").WithMessage("سختی باید بین ۱ تا ۵ باشد")
	}

	//get is exit this lesson id
	fmt.Println("98sjy7")
	return m.manipulationRepo.AddLessonPending(lesson, userId)

}

func (m Manipulation) EditPendingMajor(lesson entity.PendingMajor, userId int) error {

	return m.manipulationRepo.AddMajorPending(lesson, userId)

}
