package repository

import (
	"context"
	"sync"

	"github.com/maehiyu/tollo/internal/chatservice/domain/chat"
)

type ChatRepositoryMock struct {
	mu       sync.RWMutex
	messages map[string][]*chat.Message
}

func NewChatRepositoryMock() *ChatRepositoryMock {
	return &ChatRepositoryMock{
		messages: make(map[string][]*chat.Message),
	}
}

func (m *ChatRepositoryMock) SaveMessage(ctx context.Context, message *chat.Message) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	m.messages[message.ChatID] = append(m.messages[message.ChatID], message)
	return nil
}

func (m *ChatRepositoryMock) FindMessagesByChatID(ctx context.Context, chatID string) ([]*chat.Message, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()

	if messages, ok := m.messages[chatID]; ok {
		return messages, nil
	}

	return []*chat.Message{}, nil
}
