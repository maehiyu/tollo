package chat

import "context"

type ChatRepository interface {
	SaveMessage(ctx context.Context, message *Message) error
	FindMessagesByChatID(ctx context.Context, chatID string) ([]*Message, error)
}
