package repository

import (
	"context"
	"sync"
	"time"

	"github.com/google/uuid"

	"github.com/maehiyu/tollo/internal/chatservice/domain/chat"
)

type ChatRepositoryMock struct {
	mu       sync.RWMutex
	messages map[string][]*chat.Message
	chats    map[string]*chat.Chat
}

func NewChatRepositoryMock() *ChatRepositoryMock {
	repo := &ChatRepositoryMock{
		messages: make(map[string][]*chat.Message),
		chats:    make(map[string]*chat.Chat),
	}

	// Add some mock chats
	chat1ID := uuid.New().String()
	chat2ID := uuid.New().String()

	repo.chats[chat1ID] = &chat.Chat{
		ID:                 chat1ID,
		GeneralUserID:      "user1-uuid", // Sample UUID
		ProfessionalUserID: "user2-uuid", // Sample UUID
		CreatedAt:          time.Now().Add(-2 * time.Hour),
		UpdatedAt:          time.Now().Add(-1 * time.Hour),
		LatestMessage: &chat.Message{
			ID:       uuid.New().String(),
			ChatID:   chat1ID,
			SenderID: "user1-uuid",
			SentAt:   time.Now().Add(-30 * time.Minute),
			Payload:  &chat.StandardMessage{Content: "Hello World from chat1"},
		},
	}

	repo.chats[chat2ID] = &chat.Chat{
		ID:                 chat2ID,
		GeneralUserID:      "user3-uuid", // Sample UUID
		ProfessionalUserID: "user1-uuid", // Sample UUID
		CreatedAt:          time.Now().Add(-1 * time.Hour),
		UpdatedAt:          time.Now().Add(-30 * time.Minute),
		LatestMessage: &chat.Message{
			ID:       uuid.New().String(),
			ChatID:   chat2ID,
			SenderID: "user3-uuid",
			SentAt:   time.Now().Add(-5 * time.Minute),
			Payload:  &chat.QuestionMessage{Content: "How are you?", Tags: []string{"greeting"}},
		},
	}

	return repo
}

func (m *ChatRepositoryMock) SaveMessage(ctx context.Context, message *chat.Message) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	m.messages[message.ChatID] = append(m.messages[message.ChatID], message)
	return nil
}

func (m *ChatRepositoryMock) Create(ctx context.Context, newChat *chat.Chat) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	m.chats[newChat.ID] = newChat
	return nil
}

func (m *ChatRepositoryMock) GetMessagesByChatID(ctx context.Context, chatID string) ([]*chat.Message, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()

	if messages, ok := m.messages[chatID]; ok {
		return messages, nil
	}

	return []*chat.Message{}, nil
}

func (m *ChatRepositoryMock) GetChatsByUserID(ctx context.Context, userID string) ([]*chat.Chat, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()

	var userChats []*chat.Chat
	for _, c := range m.chats {
		if c.GeneralUserID == userID || c.ProfessionalUserID == userID {
			userChats = append(userChats, c)
		}
	}
	return userChats, nil
}
