package academicService

import (
	"fmt"
	"ostadbun/entity"
)

func (s Service) MultiDepend(id int, target string) (entity.MultiDependSlice, error) {
	data, err := s.academicRepo.MultiDepending(id, fmt.Sprintf("%s_id", target))
	if err != nil || data == nil {
		// TODO log here
		return entity.MultiDependSlice{}, err
	}

	fmt.Println(id, target, data)

	var multi = entity.MultiDependSlice{
		Major:      []entity.Major{},
		Lessons:    []entity.Lesson{},
		Professor:  []entity.Professor{},
		University: []entity.University{},
	}

	if *data != nil {
		for _, b := range *data {
			multi.Professor = append(multi.Professor, entity.Professor{Id: b.ProfessorId})
			multi.Lessons = append(multi.Lessons, entity.Lesson{Id: b.LessonId})
			multi.Major = append(multi.Major, entity.Major{Id: b.MajorId})
			multi.University = append(multi.University, entity.University{Id: b.UniversityId})
		}
	}

	fmt.Println(multi)
	newMulti := s.creation(multi)

	return newMulti, nil

}

func (s Service) creation(data entity.MultiDependSlice) entity.MultiDependSlice {
	var newData entity.MultiDependSlice

	for _, a := range data.Major {
		w, e := s.MajorGetForMulti(a.Id)
		if e != nil {
			fmt.Println("error major:", e)

			continue
		}
		newData.Major = append(newData.Major, *w)
	}

	for _, a := range data.Lessons {
		w, e := s.LessonGetForMulti(a.Id)
		if e != nil {
			fmt.Println("error major:", e)

			continue
		}
		newData.Lessons = append(newData.Lessons, *w)
	}

	for _, a := range data.University {
		w, e := s.UniversityGetForMulti(a.Id)
		if e != nil {
			fmt.Println("error major:", e)

			continue
		}
		newData.University = append(newData.University, *w)
	}

	for _, a := range data.Professor {
		w, e := s.ProfessorGetForMulti(a.Id)
		if e != nil {
			fmt.Println("error major:", e)

			continue
		}
		newData.Professor = append(newData.Professor, *w)
	}

	return newData
}

func (s Service) UniversityGetForMulti(id int) (*entity.University, error) {
	data, err := s.academicRepo.UniversityGet(id)

	if err != nil {
		// TODO log here
		fmt.Println(err)
	}
	return data, err
}

func (s Service) ProfessorGetForMulti(id int) (*entity.Professor, error) {
	data, err := s.academicRepo.ProfessorGet(id)
	if err != nil {
		// TODO log here
		fmt.Println(err)
	}

	return data, err
}

func (s Service) MajorGetForMulti(id int) (*entity.Major, error) {
	data, err := s.academicRepo.MajorGet(id)
	if err != nil {
		// TODO log here
		fmt.Println(err)
	}
	return data, err
}

func (s Service) LessonGetForMulti(id int) (*entity.Lesson, error) {
	data, err := s.academicRepo.LessonGet(id)
	if err != nil {
		// TODO log here
		fmt.Println(err)
	}
	return data, err
}
