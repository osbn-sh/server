package academicService

import (
	"fmt"
	"ostadbun/entity"
)

// this service need to optimize
func (s Service) MultiDepend(id int, target string) (entity.MultiDependSlice, error) {
	data, err := s.academicRepo.MultiDepending(id, fmt.Sprintf("%s_id", target))
	if err != nil || data == nil {
		// TODO log here
		return entity.MultiDependSlice{}, err
	}

	var multi = entity.MultiDependSlice{
		Major:      []entity.Major{},
		Lessons:    []entity.Lesson{},
		Professor:  []entity.Professor{},
		University: []entity.University{},
	}

	// map برای جلوگیری از اضافه شدن آیدی تکراری
	seenMajor := map[int]bool{}
	seenLesson := map[int]bool{}
	seenProfessor := map[int]bool{}
	seenUniversity := map[int]bool{}

	if *data != nil {
		for _, b := range *data {
			if !seenProfessor[b.ProfessorId] {
				seenProfessor[b.ProfessorId] = true
				multi.Professor = append(multi.Professor, entity.Professor{Id: b.ProfessorId})
			}
			if !seenLesson[b.LessonId] {
				seenLesson[b.LessonId] = true
				multi.Lessons = append(multi.Lessons, entity.Lesson{Id: b.LessonId})
			}
			if !seenMajor[b.MajorId] {
				seenMajor[b.MajorId] = true
				multi.Major = append(multi.Major, entity.Major{Id: b.MajorId})
			}
			if !seenUniversity[b.UniversityId] {
				seenUniversity[b.UniversityId] = true
				multi.University = append(multi.University, entity.University{Id: b.UniversityId})
			}
		}
	}

	newMulti := s.creation(multi, target)

	return newMulti, nil
}
func (s Service) creation(data entity.MultiDependSlice, target string) entity.MultiDependSlice {
	var newData entity.MultiDependSlice

	if target != "major" {
		for _, a := range data.Major {
			w, e := s.MajorGetForMulti(a.Id)
			if e != nil {
				fmt.Println("error major:", e)

				continue
			}
			newData.Major = append(newData.Major, *w)
		}
	}
	if target != "lesson" {
		for _, a := range data.Lessons {
			w, e := s.LessonGetForMulti(a.Id)
			if e != nil {
				fmt.Println("error major:", e)

				continue
			}
			newData.Lessons = append(newData.Lessons, *w)
		}
	}

	if target != "university" {
		for _, a := range data.University {
			w, e := s.UniversityGetForMulti(a.Id)
			if e != nil {
				fmt.Println("error major:", e)

				continue
			}
			newData.University = append(newData.University, *w)
		}
	}

	if target != "professor" {
		for _, a := range data.Professor {
			w, e := s.ProfessorGetForMulti(a.Id)
			if e != nil {
				fmt.Println("error major:", e)

				continue
			}
			newData.Professor = append(newData.Professor, *w)
		}
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
