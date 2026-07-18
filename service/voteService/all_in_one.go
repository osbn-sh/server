package voteService

import (
	"ostadbun/entity"
	"ostadbun/param/voteparam"
)

func (s Service) RemoveRate(id int) error {
	return s.voteRepo.RemoveVote(id)
}

func (s Service) RemoveOption(id int) error {
	return s.voteRepo.RemoveOption(id)
}

func (s Service) AddRate(userID int, data voteparam.Vote) error {
	return s.voteRepo.AddRate(userID, data)
}

func (s Service) AddOption(data voteparam.Option) error {
	return s.voteRepo.AddOption(data)
}

func (s Service) UpdateRate(rateId, newRate int) error {
	return s.voteRepo.ChangeRate(rateId, newRate)
}

func (s Service) UpdateOption(optionID int, data voteparam.Option) error {
	return s.voteRepo.ChangeOption(optionID, data)
}

func (s Service) Get(data voteparam.Vote) ([]entity.OptionVoteResult, error) {
	return s.voteRepo.CalcVotesByOption(data.Target, data.TargetID)
}

func (s Service) GetMyRate(userID int, data voteparam.Vote) ([]entity.MyVote, error) {
	return s.voteRepo.GetMyRates(data.Target, userID, data.TargetID)
}

func (s Service) GetOptions() ([]voteparam.Option, error) {
	return s.voteRepo.GetOption()
}
