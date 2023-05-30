package model

import (
	"time"

	"github.com/Red-Sock/gitm8/internal/service/domain"
)

type PingPayload struct {
	Zen    string `json:"zen"`
	HookId int    `json:"hook_id"`
	Hook   struct {
		Type   string   `json:"type"`
		Id     int      `json:"id"`
		Name   string   `json:"name"`
		Active bool     `json:"active"`
		Events []string `json:"events"`
		Config struct {
			ContentType string `json:"content_type"`
			InsecureSsl string `json:"insecure_ssl"`
			Url         string `json:"url"`
		} `json:"config"`
		UpdatedAt     time.Time `json:"updated_at"`
		CreatedAt     time.Time `json:"created_at"`
		Url           string    `json:"url"`
		TestUrl       string    `json:"test_url"`
		PingUrl       string    `json:"ping_url"`
		DeliveriesUrl string    `json:"deliveries_url"`
		LastResponse  struct {
			Code    interface{} `json:"code"`
			Status  string      `json:"status"`
			Message interface{} `json:"message"`
		} `json:"last_response"`
	} `json:"hook"`
	Repository Repo `json:"repository"`
	Sender     User `json:"sender"`
}

func (p *PingPayload) GetProject() domain.Project {
	return p.Repository.ToDomain()
}

func (p *PingPayload) GetAuthor() domain.Author {
	return p.Sender.ToDomain()
}

func (p *PingPayload) GetSrcBranch() domain.Branch {
	return domain.Branch{}
}

func (p *PingPayload) GetPullRequest() domain.PullRequestPayload {
	return domain.PullRequestPayload{}
}

func (p *PingPayload) GetCommits() []domain.Commit {
	return nil
}

func (p *PingPayload) GetEventType() domain.EventType {
	return domain.Ping
}
