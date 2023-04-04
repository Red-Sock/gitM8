package git

import (
	"gitM8/internal/clients/git/github"
	"gitM8/internal/service/domain"
	"gitM8/internal/service/interfaces"
)

func GetClient(repoType domain.RepoType, URL, token string) (interfaces.Git, error) {
	switch repoType {
	case domain.RepoTypeGithub:
		return github.New(token), nil
	default:
		return nil, domain.ErrUnknownRepoType
	}
}
