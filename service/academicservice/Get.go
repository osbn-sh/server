package academicservice

import "ostadbun/entity"

func (s Service) UniversityGet(id int) (*entity.University, error) {
	data, err := s.academicRepo.UniversityGet(id)
	if err != nil {
		// TODO log here
	}
	return data, err
}

func (s Service) ProfessorGet(id int) (*entity.Professor, error) {
	data, err := s.academicRepo.ProfessorGet(id)
	if err != nil {
		// TODO log here
	}
	return data, err
}

func (s Service) MajorGet(id int) (*entity.Major, error) {
	data, err := s.academicRepo.MajorGet(id)
	if err != nil {
		// TODO log here
	}
	return data, err
}

func (s Service) LessonGet(id int) (*entity.Lesson, error) {
	data, err := s.academicRepo.LessonGet(id)
	if err != nil {
		// TODO log here
	}
	return data, err
}
