package git

import (
	"github.com/Red-Sock/gitm8/internal/clients/git/github"
	"github.com/Red-Sock/gitm8/internal/service/domain"
	"github.com/Red-Sock/gitm8/internal/service/interfaces"
)

func GetClient(repoType domain.RepoType, URL, token string) (interfaces.Git, error) {
	switch repoType {
	case domain.RepoTypeGithub:
		return github.New(token), nil
	default:
		return nil, domain.ErrUnknownRepoType
	}
}
