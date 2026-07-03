package studentService

import (
	academicRepository "ostadbun/repository/postgres/studentRepository"
)

type Service struct {
	academicRepo academicRepository.DB
}

func New(academicRepo academicRepository.DB) Service {
	return Service{academicRepo: academicRepo}
}
