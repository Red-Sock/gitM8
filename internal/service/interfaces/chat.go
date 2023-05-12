package interfaces

type Chat interface {
	Send(chatID uint64, message string)
}
