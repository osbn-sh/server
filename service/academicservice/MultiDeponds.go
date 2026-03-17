package academicservice

import (
	"fmt"
	"ostadbun/entity"
)

func (s Service) MultiDepend(id int, target string) (entity.MultiDepondMap, error) {

	//major
	//lesson
	//university
	//professor

	data, err := s.academicRepo.MultiDepending(id, fmt.Sprintf("%s_id", target))
	if err != nil || data == nil {
		// TODO log here
	}

	fmt.Println(id, target, data)

	var multi = entity.MultiDepondMap{
		Major:      make(map[int]entity.Major),
		Lessons:    make(map[int]entity.Lesson),
		Professor:  make(map[int]entity.Professor),
		University: make(map[int]entity.University),
	}

	if *data != nil {
		for _, b := range *data {
			multi.Professor[b.ProfessorId] = entity.Professor{}
			multi.University[b.UniversityId] = entity.University{}
			multi.Major[b.MajorId] = entity.Major{}
			multi.Lessons[b.LessonId] = entity.Lesson{}
		}
	}

	s.creation(&multi)

	return multi, nil

}

func (s Service) creation(data *entity.MultiDepondMap) {
	for a, _ := range data.Major {
		w, e := s.MajorGetForMulti(a)
		if e != nil {
			fmt.Println("error major:", e)

			continue
		}
		data.Major[a] = *w
	}

	for a, _ := range data.Lessons {
		w, e := s.LessonGetForMulti(a)
		if e != nil {
			fmt.Println("error lesson:", e)

			continue
		}
		data.Lessons[a] = *w
	}

	for a, _ := range data.Professor {
		w, e := s.ProfessorGetForMulti(a)
		if e != nil {
			fmt.Println("error professor:", e)

			continue
		}
		data.Professor[a] = *w
	}

	for a, _ := range data.University {
		w, e := s.UniversityGetForMulti(a)
		if e != nil {
			fmt.Println("error university:", e)

			continue
		}
		data.University[a] = *w
	}
}

func (s Service) UniversityGetForMulti(id int) (*entity.University, error) {
	data, err := s.academicRepo.UniversityGet(id)
	count, err := s.academicRepo.UserCountUniversity(id)

	if err != nil {
		// TODO log here
		fmt.Println(err)
	}
	if data != nil {
		data.UsersCount = count
	}

	//fmt.Println(multi)
	return data, err
}

func (s Service) ProfessorGetForMulti(id int) (*entity.Professor, error) {
	data, err := s.academicRepo.ProfessorGet(id)
	count, err := s.academicRepo.UserCountProfessor(id)
	if err != nil {
		// TODO log here
		fmt.Println(err)
	}
	if data != nil {

		data.UsersCount = len(count)
	}
	return data, err
}

func (s Service) MajorGetForMulti(id int) (*entity.Major, error) {

	data, err := s.academicRepo.MajorGet(id)
	count, err := s.academicRepo.UserCountMajor(id)
	if err != nil {
		// TODO log here
		fmt.Println(err)
	}

	if data != nil {
		data.UsersCount = count
	}
	return data, err
}

func (s Service) LessonGetForMulti(id int) (*entity.Lesson, error) {
	data, err := s.academicRepo.LessonGet(id)
	count, err := s.academicRepo.UserCountLesson(id)

	if err != nil {
		// TODO log here
		fmt.Println(err)
	}
	if data != nil {
		data.UsersCount = len(count)
	}
	return data, err
}
