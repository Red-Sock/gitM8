package git

import (
	"github.com/Red-Sock/gitm8/internal/clients/git/github"
	"github.com/Red-Sock/gitm8/internal/service/domain/webhook"
	"github.com/Red-Sock/gitm8/internal/service/interfaces"
)

func GetClient(repoType webhook.RepoType, URL, token string) (interfaces.Git, error) {
	switch repoType {
	case webhook.RepoTypeGithub:
		return github.New(token), nil
	default:
		return nil, webhook.ErrUnknownRepoType
	}
}
