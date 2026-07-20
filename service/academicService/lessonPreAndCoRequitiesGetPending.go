package academicService

import (
	"ostadbun/entity"
)

func (s Service) GetPendingLessonCheck() (entity.OutParam, error) {
	dataPre, errPre := s.academicRepo.GetPendingLessonPreRequisites()

	dataCo, errCo := s.academicRepo.GetPendingLessonCoRequisites()

	if errPre != nil || errCo != nil {
		return entity.OutParam{}, errPre
	}

	data := entity.OutParam{
		PreRequites: dataPre,
		CoRequites:  dataCo,
	}

	return data, nil

}
