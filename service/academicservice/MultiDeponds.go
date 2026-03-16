package academicservice

import (
	"fmt"
	"ostadbun/entity"
)

func (s Service) MultiDependLesson(id int) (entity.MultiDepondMap, error) {
	return s.MultiDepend(id, "lesson")
}

func (s Service) MultiDependMajor(id int) (entity.MultiDepondMap, error) {
	return s.MultiDepend(id, "major")
}

func (s Service) MultiDependProfessor(id int) (entity.MultiDepondMap, error) {
	return s.MultiDepend(id, "professor")
}

func (s Service) MultiDependUniversity(id int) (entity.MultiDepondMap, error) {
	return s.MultiDepend(id, "university")
}
func (s Service) MultiDepend(id int, target string) (entity.MultiDepondMap, error) {

	//major
	//lesson
	//university
	//professor

	data, err := s.academicRepo.MultiDepending(id, fmt.Sprintf("%s_id", target))
	if err != nil || data == nil {
		// TODO log here
	}

	fmt.Println(data)

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
		w, e := s.MajorGet(a)
		if e != nil {
			fmt.Println("error major:", e)

			continue
		}
		data.Major[a] = *w
	}

	for a, _ := range data.Lessons {
		w, e := s.LessonGet(a)
		if e != nil {
			fmt.Println("error lesson:", e)

			continue
		}
		data.Lessons[a] = *w
	}

	for a, _ := range data.Professor {
		w, e := s.ProfessorGet(a)
		if e != nil {
			fmt.Println("error professor:", e)

			continue
		}
		data.Professor[a] = *w
	}

	for a, _ := range data.University {
		w, e := s.UniversityGet(a)
		if e != nil {
			fmt.Println("error university:", e)

			continue
		}
		data.University[a] = *w
	}
}
