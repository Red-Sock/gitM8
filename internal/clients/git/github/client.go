package github

import (
	"context"
	"errors"
	"net/http"

	cli "github.com/google/go-github/v50/github"
	"golang.org/x/oauth2"

	gitErrors "gitM8/internal/clients/git/errors"
	"gitM8/internal/service/domain"
)

func New(token string) *GitClient {
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	tc := oauth2.NewClient(ctx, ts)

	client := cli.NewClient(tc)
	return &GitClient{
		cli: client,
	}
}

type GitClient struct {
	cli *cli.Client
}

func (g *GitClient) GetCurrentUser(ctx context.Context) (domain.TgUser, error) {
	usr, rsp, err := g.cli.Users.Get(ctx, "")
	if err != nil {
		if rsp.StatusCode == http.StatusUnauthorized {
			return domain.TgUser{}, errors.Join(err, gitErrors.ErrUnauthorized)
		}

		return domain.TgUser{}, errors.Join(err, gitErrors.ErrCouldNotFindCurrentUser)
	}

	var out domain.TgUser
	if usr.ID == nil {
		return domain.TgUser{}, errors.Join(gitErrors.ErrInvalidResponseData, errors.New("no user id"))
	}
	out.GitId = uint64(*usr.ID)

	if usr.Name == nil {
		return domain.TgUser{}, errors.Join(gitErrors.ErrInvalidResponseData, errors.New("no username"))
	}
	out.GitUsername = *usr.Name

	return out, nil
}
