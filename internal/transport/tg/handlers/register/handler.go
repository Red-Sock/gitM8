package register

import (
	tginterfaces "github.com/AlexSkilled/go_tg/interfaces"
	"github.com/AlexSkilled/go_tg/model"
)

const Command = "register"

type MainMenu struct {
}

func New() *MainMenu {
	return &MainMenu{}
}

func (m MainMenu) Handle(in *model.MessageIn, out tginterfaces.Sender) {

}

func (m MainMenu) Dump(chatId int64) {

}
