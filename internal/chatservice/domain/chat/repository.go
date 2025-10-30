package chat

import "context"

type ChatRepository interface {
	SaveMessage(ctx context.Context, message *Message) error
	GetChatsByUserID(ctx context.Context, userID string) ([]*Chat, error)
	GetMessagesByChatID(ctx context.Context, chatID string) ([]*Message, error)
}
