package votehandler

import (
	"ostadbun/service/userservice"
	"ostadbun/service/voteService"
)

type Handler struct {
	userSvc userservice.User
	voteSvc voteService.Service
}

func New(userSvc userservice.User, voteSvc voteService.Service) Handler {
	return Handler{
		userSvc: userSvc,
		voteSvc: voteSvc,
	}
}
