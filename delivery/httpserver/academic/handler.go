package academic

import (
	academicservice "ostadbun/service/academicService"
)

type Handler struct {
	academicService academicservice.Service
}

func New(academicService academicservice.Service) Handler {
	return Handler{
		academicService: academicService,
	}
}
