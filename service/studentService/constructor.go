package studentService

import (
	academicRepository "ostadbun/repository/postgres/studentRepository"
	"ostadbun/service/academicService"
	"ostadbun/service/activityService"
)

type Service struct {
	academicRepo academicRepository.DB
	academicSvc  academicService.Service
	activity     activityService.Activity
}

func New(academicRepo academicRepository.DB, activity activityService.Activity, academicSvc academicService.Service) Service {
	return Service{academicRepo: academicRepo, activity: activity, academicSvc: academicSvc}
}
