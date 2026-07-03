package backgroundJobs

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"ostadbun/service/githubcheckingversionservice"
	"sync"
	"time"
)

func (b *BackJob) Github(e githubcheckingversionservice.GithubCheckingVersionService) error {

	_, err := b.cron.AddFunc("@weekly", func() {
		do(e)
	})

	return err
}

type Release struct {
	TagName string `json:"tag_name"`
	Name    string `json:"name"`
}

func do(e githubcheckingversionservice.GithubCheckingVersionService) {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	owner := os.Getenv("GITHUB_OWNER")
	repoClient := os.Getenv("GITHUB_REPO_CLIENT")
	repoServer := os.Getenv("GITHUB_REPO_SERVER")

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		response, err := httpRequest(owner, repoClient)
		if err != nil {
			log.Printf("github request for client sync failed: %v", err)

			return
		}

		errRds := e.SetClientVersion(ctx, response.Name)
		if errRds != nil {
			log.Printf("github redis for client sync failed: %v", err)

			return
		}
	}()

	go func() {
		defer wg.Done()
		response, err := httpRequest(owner, repoServer)
		if err != nil {
			log.Printf("github request for client sync failed: %v", err)

			return
		}

		errRds := e.SetServerVersion(ctx, response.Name)
		if errRds != nil {
			log.Printf("github redis for client sync failed: %v", err)

			return
		}
	}()

	wg.Wait()
}

func httpRequest(owner, repo string) (Release, error) {

	url := fmt.Sprintf("https://api.github.com/repos/%s/%s/releases/latest", owner, repo)

	fmt.Println("🪏 background job starting")
	req, err := http.NewRequest(
		http.MethodGet,
		url,
		nil,
	)
	if err != nil {
		return Release{}, err
	}

	req.Header.Set("Accept", "application/vnd.github+json")

	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	resp, err := client.Do(req)
	if err != nil {
		return Release{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return Release{}, fmt.Errorf("github returned %d", resp.StatusCode)
	}

	var release Release
	err = json.NewDecoder(resp.Body).Decode(&release)
	if err != nil {
		return Release{}, err
	}

	return release, nil
}
