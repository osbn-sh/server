package homehandler

import "ostadbun/service/githubcheckingversionservice"

type Handler struct {
	GithubCheckingVersionService githubcheckingversionservice.GithubCheckingVersionService
}

func New(GithubCheckingVersionService githubcheckingversionservice.GithubCheckingVersionService) Handler {
	return Handler{
		GithubCheckingVersionService: GithubCheckingVersionService,
	}
}
