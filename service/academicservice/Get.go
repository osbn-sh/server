package academicservice

import (
	"fmt"
	"ostadbun/entity"
)

func (s Service) UniversityGet(id int) (*entity.University, error) {
	data, err := s.academicRepo.UniversityGet(id)
	count, err := s.academicRepo.UserCountUniversity(id)
	multi, err := s.MultiDepend(id, "university")

	if err != nil {
		// TODO log here
		fmt.Println(err)
	}
	if data != nil {
		data.UsersCount = count
		data.Relationships = &multi

	}

	//fmt.Println(multi)
	return data, err
}

func (s Service) ProfessorGet(id int) (*entity.Professor, error) {
	data, err := s.academicRepo.ProfessorGet(id)
	count, err := s.academicRepo.UserCountProfessor(id)
	multi, err := s.MultiDepend(id, "professor")
	if err != nil {
		// TODO log here
		fmt.Println(err)
	}
	if data != nil {

		data.UsersCount = len(count)
		data.Relationships = &multi
	}
	return data, err
}

func (s Service) MajorGet(id int) (*entity.Major, error) {

	data, err := s.academicRepo.MajorGet(id)
	count, err := s.academicRepo.UserCountMajor(id)
	multi, err := s.MultiDepend(id, "major")
	if err != nil {
		// TODO log here
		fmt.Println(err)
	}

	if data != nil {
		data.UsersCount = count
		data.Relationships = &multi
	}
	return data, err
}

func (s Service) LessonGet(id int) (*entity.Lesson, error) {
	data, err := s.academicRepo.LessonGet(id)
	count, err := s.academicRepo.UserCountLesson(id)
	multi, err := s.MultiDepend(id, "lesson")
	PreRequ, err := s.LessonPreReq(id)
	CoRequ, err := s.LessonCoReq(id)

	if err != nil {
		// TODO log here
		fmt.Println(err)
	}

	if data != nil {
		data.UsersCount = len(count)
		data.Relationships = &multi

		if PreRequ != nil {
			data.PreRequites = PreRequ
		}

		if CoRequ != nil {
			data.CoRequites = CoRequ
		}
	}

	return data, err
}
