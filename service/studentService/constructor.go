package studentService

import (
	academicRepository "ostadbun/repository/postgres/studentRepository"
	"ostadbun/service/activityService"
)

type Service struct {
	academicRepo academicRepository.DB
	activity     activityService.Activity
}

func New(academicRepo academicRepository.DB, activity activityService.Activity) Service {
	return Service{academicRepo: academicRepo, activity: activity}
}
