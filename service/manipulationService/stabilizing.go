package manipulationService

func (m Manipulation) StabilizeLesson(lessonID int) error {

	return m.manipulationRepo.StabilizeLesson(lessonID)

}

func (m Manipulation) StabilizeProfessor(professorID int) error {

	return m.manipulationRepo.StabilizeProfessor(professorID)

}

func (m Manipulation) StabilizeMajor(majorID int) error {

	return m.manipulationRepo.StabilizeMajor(majorID)

}

func (m Manipulation) StabilizeUniversity(universityID int) error {

	return m.manipulationRepo.StabilizeUniversity(universityID)

}
