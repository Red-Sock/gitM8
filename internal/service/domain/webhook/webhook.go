package webhook

type Request struct {
	Src     RepoType
	Type    Type
	Payload []byte
}

type Type int

const (
	Invalid Type = iota
	Push
	PullRequest
)

func (w *Type) ParseGithub(in string) {
	*w, _ = githubEventsToDomain[in]
}

var githubEventsToDomain = map[string]Type{
	"push":         Push,
	"pull_request": PullRequest,
}
