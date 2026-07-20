package academicService

import (
	"ostadbun/entity"
)

func (s Service) GetPendingLessonPreReq(id int) (*[]entity.Lesson, error) {
	//data, err := s.academicRepo.GetPendingLessonPreRequisites(id)
	//if err != nil {
	//	return nil, richerror.New("academicRepository-LessonPreReq").WithErr(err)
	//}

	//lessons, errL := s.academicRepo.LessonGetMany(data...)
	//
	//if errL != nil {
	//	return nil, richerror.New("academicRepository-LessonPreReq").WithErr(errL).WithMessage("error getting lessons on pre requities")
	//
	//}

	//return lessons, err
	panic("")
}

func (s Service) GetPendingLessonCoReq(id int) (*[]entity.Lesson, error) {
	//data, err := s.academicRepo.GetPendingLessonCoRequisites(id)
	//if err != nil {
	//	return nil, richerror.New("academicRepository-LessonCoReq").WithErr(err)
	//}
	//
	//lessons, errL := s.academicRepo.LessonGetMany(data...)
	//
	//if errL != nil {
	//	return nil, richerror.New("academicRepository-LessonCoReq").WithErr(errL).WithMessage("error getting lessons on Co requities")
	//
	//}
	//
	//return lessons, err
	panic("")
}
