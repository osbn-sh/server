package manipulationService

import (
	"fmt"
	"ostadbun/entity"
	"ostadbun/pkg/richerror"
)

func (m Manipulation) EditPendingUniversity(lesson entity.PendingUniversity, userId int) error {

	pl, _ := m.academic.UniversityGet(fmt.Sprintf("%d", *lesson.TargetId))

	OkLogic := pl != nil

	if !OkLogic {
		return richerror.New("addpendinglesson.service").WithMessage(fmt.Sprintf("همچین چیزی وجود ندارد %d", *lesson.TargetId))
	}

	return m.manipulationRepo.AddUniversityPending(lesson, userId)

}

func (m Manipulation) EditPendingProfessor(lesson entity.PendingProfessor, userId int) error {

	pl, _ := m.academic.ProfessorGet(fmt.Sprintf("%d", *lesson.TargetId))

	OkLogic := pl != nil

	if !OkLogic {
		return richerror.New("addpendinglesson.service").WithMessage(fmt.Sprintf("همچین چیزی وجود ندارد %d", *lesson.TargetId))
	}

	return m.manipulationRepo.AddProfessorPending(lesson, userId)

}

func (m Manipulation) EditPendingLesson(lesson entity.PendingLesson, userId int) error {

	if lesson.Difficulty > 5 || lesson.Difficulty < 1 {
		return richerror.New("addpendinglesson.service").WithMessage("سختی باید بین ۱ تا ۵ باشد")
	}

	pl, _ := m.academic.LessonGet(fmt.Sprintf("%d", *lesson.TargetId))

	OkLogic := pl != nil

	if !OkLogic {
		return richerror.New("addpendinglesson.service").WithMessage(fmt.Sprintf("همچین چیزی وجود ندارد %d", *lesson.TargetId))
	}
	return m.manipulationRepo.AddLessonPending(lesson, userId)

}

func (m Manipulation) EditPendingMajor(lesson entity.PendingMajor, userId int) error {

	pl, _ := m.academic.MajorGet(fmt.Sprintf("%d", *lesson.TargetId))

	OkLogic := pl != nil

	if !OkLogic {
		return richerror.New("addpendinglesson.service").WithMessage(fmt.Sprintf("همچین چیزی وجود ندارد %d", *lesson.TargetId))
	}

	return m.manipulationRepo.AddMajorPending(lesson, userId)

}
