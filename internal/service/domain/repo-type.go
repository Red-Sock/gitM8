package domain

import (
	"strings"

	"github.com/pkg/errors"
)

var ErrUnknownRepoType = errors.New("unknown repo type")

type RepoType int

const (
	RepoTypeInvalid RepoType = iota
	RepoTypeGithub
)

const Github = "github"

func (r *RepoType) String() string {
	return ""
}

func (r *RepoType) SetType(in string) error {

	for k, v := range repoTypeMap {
		if strings.Contains(in, k) {
			*r = v
			return nil
		}
	}

	return errors.Wrapf(ErrUnknownRepoType, "no such repo type - %s", in)
}

func (r *RepoType) GetHelpMessage() string {
	switch *r {
	case RepoTypeGithub:
		return "https://github.com/settings/tokens/new?description=RedSock%20Webhook%20telegram%20plugin&scopes=read:user"
	default:
		return "unknown git source"
	}
}

var repoTypeMap = map[string]RepoType{
	Github: RepoTypeGithub,
}
