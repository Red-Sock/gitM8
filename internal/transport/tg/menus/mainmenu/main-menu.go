package mainmenu

import (
	"github.com/AlexSkilled/go_tg/interfaces"
	"github.com/AlexSkilled/go_tg/model/menu"

	"gitM8/internal/transport/tg/handlers/register"
)

const Command = "menu"

func NewMainMenu() interfaces.Menu {
	m := menu.NewSimple(Command, "Main menu")

	m.AddButton(register.Command, "Register")

	return m
}
