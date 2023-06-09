package github

import (
	"context"
	"errors"
	"net/http"

	cli "github.com/google/go-github/v50/github"
	"golang.org/x/oauth2"

	gitErrors "github.com/Red-Sock/gitm8/internal/clients/git/errors"
	"github.com/Red-Sock/gitm8/internal/service/domain"
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

	if usr.Name == nil {
		return domain.TgUser{}, errors.Join(gitErrors.ErrInvalidResponseData, errors.New("no username"))
	}

	return out, nil
}
