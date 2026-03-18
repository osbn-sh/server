package studentService

import (
	"ostadbun/param/studentparam"
)

func (s Service) Remove(id, userID int) (bool, error) {

	return s.academicRepo.RemovePass(id, userID)

}

func (s Service) Add(userID int, student studentparam.StudentPassDetail) error {

	err := s.academicRepo.AddPass(userID, student)

	return err
}
