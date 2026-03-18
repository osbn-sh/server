package student

import (
	academicservice "ostadbun/service/academicService"
	"ostadbun/service/studentService"
	"ostadbun/service/userservice"
)

type Handler struct {
	academicService academicservice.Service
	userService     userservice.User
	studentService  studentService.Service
}

func New(academicService academicservice.Service, studentService studentService.Service, userService userservice.User) Handler {
	return Handler{
		academicService: academicService,
		userService:     userService,
		studentService:  studentService,
	}
}
