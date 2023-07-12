package tg_message_constructor

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

	"github.com/Red-Sock/gitm8/internal/service/domain"
	"github.com/Red-Sock/gitm8/internal/transport/tg/assets"
)

type releaseConstructor struct {
}

func (m *MessageConstructor) extractRelease(payload domain.Payload) (string, []tgbotapi.MessageEntity, error) {

	asset := assets.Dunno
	phrase := "== I don't know what happened =="
	switch payload.GetAction() {
	case domain.ActionCreated:
		asset = assets.Package
		phrase = " has been created"
	case domain.Deleted:
		asset = assets.Delete
		phrase = " has been deleted"
	case domain.Edited:
		asset = assets.Edit
		phrase = " has been edited"
	case domain.PreReleased:
		asset = assets.YellowSquare
		phrase = " has been pre released"
	case domain.Published:
		asset = assets.Party
		phrase = " has been published!"
	case domain.Released:
		asset = assets.Party
		phrase = " has been released!"
	case domain.UnPublished:
		asset = assets.Back
		phrase = " has been unpublished :/"

	default:
		return "", nil, domain.ErrInvalidActionForEvent
	}

	constr := constructor{}

	constr.Write(asset)
	constr.Write("Release ")

	{
		release := payload.GetRelease()
		constr.WriteWithLink(release.Name, release.Link)
	}
	constr.Write(phrase)

	{
		constr.Writeln("Project: ")
		proj := payload.GetProject()
		constr.WriteWithLink(proj.Name, proj.Link)
	}
	{
		constr.Writeln("Author: ")
		author := payload.GetAuthor()
		constr.WriteWithLink(author.Name, author.Link)
	}
	return constr.String(), constr.format, nil

}
