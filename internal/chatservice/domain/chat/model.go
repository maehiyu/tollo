package chat

import "time"

type MessageContent interface {
	isMessageContent()
}

type StandardContent struct {
	Content string
}

func (c *StandardContent) isMessageContent() {}

type QuestionContent struct {
	Content string
	Tags    []string
}

func (c *QuestionContent) isMessageContent() {}

type AnswerContent struct {
	Content    string
	QuestionID string
}

func (c *AnswerContent) isMessageContent() {}

type PromotionalContent struct {
	Title     string
	Body      string
	ActionURL string
	ImageURL  string
}

func (c *PromotionalContent) isMessageContent() {}

type Chat struct {
	ID                 string
	GeneralUserID      string
	ProfessionalUserID string
	CreatedAt          time.Time
	UpdatedAt          time.Time
}

type Message struct {
	ID       string
	ChatID   string
	SenderID string
	Content  MessageContent
	SentAt   time.Time
}
