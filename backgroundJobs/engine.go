package backgroundJobs

import (
	"fmt"
	"ostadbun/service/githubcheckingversionservice"

	"github.com/robfig/cron/v3"
)

type BackJob struct {
	cron *cron.Cron
}

func New(
	GithubCheckingVersionService githubcheckingversionservice.GithubCheckingVersionService,
) *BackJob {
	b := &BackJob{
		cron: cron.New(),
	}

	if err := b.Github(GithubCheckingVersionService); err != nil {
		panic(err)
	}

	return b
}

func (b *BackJob) Stop() {
	ctx := b.cron.Stop()
	<-ctx.Done()
}

func (b *BackJob) Start() {
	fmt.Println("starting background job")
	b.cron.Start()
}
