package academicService

import (
	"ostadbun/entity"
	"strconv"
)

func (s Service) UniversityGet(target string) (*entity.University, error) {

	intTarget, errT := strconv.Atoi(target)

	var data *entity.University
	var err error

	if errT != nil {
		data, err = s.academicRepo.UniversityGetByHref(target)

	} else {
		data, err = s.academicRepo.UniversityGet(intTarget)
	}

	if data == nil || err != nil || data.Id < 1 {
		return nil, err
	}
	count, err := s.academicRepo.UserCountUniversity(data.Id)
	multi, err := s.MultiDepend(data.Id, "university")

	if err != nil {
		return nil, err
	}

	data.UsersCount = count
	data.Relationships = &multi

	return data, nil
}

func (s Service) ProfessorGet(target string) (*entity.Professor, error) {
	intTarget, errT := strconv.Atoi(target)

	var data *entity.Professor
	var err error

	if errT != nil {
		data, err = s.academicRepo.ProfessorGetByHref(target)

	} else {
		data, err = s.academicRepo.ProfessorGet(intTarget)
	}

	if data == nil || err != nil || data.Id < 1 {
		return nil, err
	}
	count, err := s.academicRepo.UserCountProfessor(data.Id)
	multi, err := s.MultiDepend(data.Id, "professor")

	if err != nil {
		return nil, err
	}

	data.UsersCount = len(count)
	data.Relationships = &multi

	return data, nil
}

func (s Service) MajorGet(target string) (*entity.Major, error) {
	intTarget, errT := strconv.Atoi(target)

	var data *entity.Major
	var err error

	if errT != nil {
		data, err = s.academicRepo.MajorGetByHref(target)

	} else {
		data, err = s.academicRepo.MajorGet(intTarget)
	}

	if data == nil || err != nil || data.Id < 1 {
		return nil, err
	}
	count, err := s.academicRepo.UserCountMajor(data.Id)
	multi, err := s.MultiDepend(data.Id, "major")

	if err != nil {
		return nil, err
	}

	data.UsersCount = count
	data.Relationships = &multi

	return data, nil
}

func (s Service) LessonGet(target string) (*entity.Lesson, error) {
	intTarget, errT := strconv.Atoi(target)

	var data *entity.Lesson
	var err error

	if errT != nil {
		data, err = s.academicRepo.LessonGetByHref(target)

	} else {
		data, err = s.academicRepo.LessonGet(intTarget)
	}

	if data == nil || err != nil || data.Id < 1 {
		return nil, err
	}

	PreRequ, err := s.LessonPreReq(data.Id)
	CoRequ, err := s.LessonCoReq(data.Id)
	count, err := s.academicRepo.UserCountLesson(data.Id)
	multi, err := s.MultiDepend(data.Id, "lesson")

	if err != nil {
		return nil, err
	}

	data.UsersCount = len(count)
	data.Relationships = &multi

	if PreRequ != nil {
		data.PreRequites = PreRequ
	}

	if CoRequ != nil {
		data.CoRequites = CoRequ
	}
	return data, nil
}
