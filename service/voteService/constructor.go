package voteService

import "ostadbun/repository/postgres/voteRepository"

type Service struct {
	voteRepo voteRepository.DB
}

func New(voteRepo voteRepository.DB) Service {
	return Service{
		voteRepo: voteRepo,
	}
}
