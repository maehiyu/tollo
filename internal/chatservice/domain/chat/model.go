package chat

import "time"

type MessagePayload interface {
	isMessagePayload()
}

type StandardMessage struct {
	Content string
}

func (c *StandardMessage) isMessagePayload() {}

type QuestionMessage struct {
	Content string
	Tags    []string
}

func (c *QuestionMessage) isMessagePayload() {}

type AnswerMessage struct {
	Content    string
	QuestionID string
}

func (c *AnswerMessage) isMessagePayload() {}

type PromotionalMessage struct {
	Title     string
	Body      string
	ActionURL string
	ImageURL  string
}

func (c *PromotionalMessage) isMessagePayload() {}

type Chat struct {
	ID                 string
	GeneralUserID      string
	ProfessionalUserID string
	CreatedAt          time.Time
	UpdatedAt          time.Time
	LatestMessage      *Message
}

type Message struct {
	ID       string
	ChatID   string
	SenderID string
	Payload  MessagePayload
	SentAt   time.Time
}
