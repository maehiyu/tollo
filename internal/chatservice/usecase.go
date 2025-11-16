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

func (u *Usecase) SendMessage(ctx context.Context, chatID, senderID string, payload chat.MessagePayload) (*chat.Message, error) {
	msg := &chat.Message{
		ID:       uuid.NewString(),
		ChatID:   chatID,
		SenderID: senderID,
		Payload:  payload,
		SentAt:   time.Now(),
	}

	if err := u.chatRepo.SaveMessage(ctx, msg); err != nil {
		return nil, err
	}

	return msg, nil
}

func (u *Usecase) GetChatsByUserID(ctx context.Context, userID string) ([]*chat.Chat, error) {
	return u.chatRepo.GetChatsByUserID(ctx, userID)
}

func (u *Usecase) GetMessagesByChatID(ctx context.Context, chatID string) ([]*chat.Message, error) {
	return u.chatRepo.GetMessagesByChatID(ctx, chatID)
}

func (u *Usecase) CreateChat(ctx context.Context, generalUserID, professionalUserID string) (*chat.Chat, error) {
	newChat := &chat.Chat{
		ID:                 uuid.NewString(),
		GeneralUserID:      generalUserID,
		ProfessionalUserID: professionalUserID,
		CreatedAt:          time.Now(),
		UpdatedAt:          time.Now(),
	}
	// The mock repository's Create method will add this to its internal map
	if err := u.chatRepo.Create(ctx, newChat); err != nil {
		return nil, err
	}
	return newChat, nil
}
