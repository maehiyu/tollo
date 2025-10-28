package chatservice

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/maehiyu/tollo/internal/chatservice/domain/chat"
)

type Usecase struct {
	chatRepo chat.ChatRepository
}

func NewUsecase(chatRepo chat.ChatRepository) *Usecase {
	return &Usecase{
		chatRepo: chatRepo,
	}
}

func (u *Usecase) SendMessage(ctx context.Context, chatID, senderID string, content chat.MessageContent) (*chat.Message, error) {
	msg := &chat.Message{
		ID:       uuid.NewString(),
		ChatID:   chatID,
		SenderID: senderID,
		Content:  content,
		SentAt:   time.Now(),
	}

	if err := u.chatRepo.SaveMessage(ctx, msg); err != nil {
		return nil, err
	}

	return msg, nil
}
