package tg_message_constructor

import (
	"strings"

	"github.com/Red-Sock/go_tg/model/response"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/mattn/go-runewidth"
	_ "github.com/mattn/go-runewidth"
)

type constructor struct {
	text   []rune
	format []tgbotapi.MessageEntity
	idx    int
}

func (c *constructor) Write(text string) {
	c.idx += runewidth.StringWidth(text)

	c.idx += strings.Count(text, "\n")
	c.text = append(c.text, []rune(text)...)
}

func (c *constructor) Writeln(text string) {
	text = "\n" + text
	c.Write(text)
}

func (c *constructor) WriteWithFormat(text, format string) {
	c.format = append(c.format, tgbotapi.MessageEntity{
		Type:   format,
		Offset: c.idx,
		Length: len(text),
	})

	c.Write(text)
}

func (c *constructor) WriteWithLink(text, url string) {
	c.format = append(c.format, tgbotapi.MessageEntity{
		Type:   response.TextLinkTextFormat,
		Offset: c.idx,
		Length: len(text),
		URL:    url,
	})

	c.Write(text)
}

func (c *constructor) WriteWithMention(text string, userID uint64) {
	c.format = append(c.format, tgbotapi.MessageEntity{
		Type:   response.MentionTextFormat,
		Offset: c.idx,
		Length: len(text),
		User: &tgbotapi.User{
			ID: int64(userID),
		},
	})

	c.Write(text)
}

func (c *constructor) String() string {
	return string(c.text)
}
