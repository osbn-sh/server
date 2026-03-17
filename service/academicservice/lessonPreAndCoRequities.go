package academicservice

import (
	"ostadbun/entity"
	"ostadbun/pkg/richerror"
)

func (s Service) LessonPreReq(id int) (*[]entity.Lesson, error) {
	data, err := s.academicRepo.LessonPreRequisites(id)
	if err != nil {
		return nil, richerror.New("academicRepository-LessonPreReq").WithErr(err)
	}

	lessons, errL := s.academicRepo.LessonGetMany(data...)

	if errL != nil {
		return nil, richerror.New("academicRepository-LessonPreReq").WithErr(errL).WithMessage("error getting lessons on pre requities")

	}
	
	return lessons, err
}

func (s Service) LessonCoReq(id int) (*[]entity.Lesson, error) {
	data, err := s.academicRepo.LessonCoRequisites(id)
	if err != nil {
		return nil, richerror.New("academicRepository-LessonCoReq").WithErr(err)
	}

	lessons, errL := s.academicRepo.LessonGetMany(data...)

	if errL != nil {
		return nil, richerror.New("academicRepository-LessonCoReq").WithErr(errL).WithMessage("error getting lessons on Co requities")

	}

	return lessons, err
}
