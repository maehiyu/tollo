package converter

import (
	"google.golang.org/protobuf/types/known/timestamppb"

	chatpb "github.com/maehiyu/tollo/gen/go/protos/chatservice"
	userpb "github.com/maehiyu/tollo/gen/go/protos/userservice"
	"github.com/maehiyu/tollo/internal/chatservice/domain/chat"
	"github.com/maehiyu/tollo/internal/userservice/domain/user"
)

func ToUserInfo(u *user.User) *userpb.UserInfo {
	if u == nil {
		return nil
	}

	userInfo := &userpb.UserInfo{
		Id:        u.ID,
		Name:      u.Name,
		Email:     string(u.Email),
		CreatedAt: timestamppb.New(u.CreatedAt),
		UpdatedAt: timestamppb.New(u.UpdatedAt),
	}

	switch p := u.Profile.(type) {
	case *user.ProfessionalProfile:
		userInfo.Profile = &userpb.UserInfo_ProfessionalProfile{
			ProfessionalProfile: &userpb.ProfessionalProfile{
				ProBadgeUrl: p.ProBadgeURL,
				Biography:   p.Biography,
			},
		}
	case *user.GeneralProfile:
		userInfo.Profile = &userpb.UserInfo_GeneralProfile{
			GeneralProfile: &userpb.GeneralProfile{
				Points:       p.Points,
				Introduction: p.Introduction,
			},
		}
	}
	return userInfo
}

// Renamed from StandardMessageProtoToDomain
func StandardPayloadProtoToDomain(p *chatpb.StandardMessage) chat.MessagePayload {
	return &chat.StandardMessage{
		Content: p.GetContent(),
	}
}

// Renamed from QuestionMessageProtoToDomain
func QuestionPayloadProtoToDomain(p *chatpb.QuestionMessage) chat.MessagePayload {
	return &chat.QuestionMessage{
		Content: p.GetContent(),
		Tags:    p.GetTags(),
	}
}

// Renamed from AnswerMessageProtoToDomain
func AnswerPayloadProtoToDomain(p *chatpb.AnswerMessage) chat.MessagePayload {
	return &chat.AnswerMessage{
		Content:    p.GetContent(),
		QuestionID: p.GetQuestionId(),
	}
}

// Renamed from PromotionalMessageProtoToDomain
func PromotionalPayloadProtoToDomain(p *chatpb.PromotionalMessage) chat.MessagePayload {
	return &chat.PromotionalMessage{
		Title:     p.GetTitle(),
		Body:      p.GetBody(),
		ActionURL: p.GetActionUrl(),
		ImageURL:  p.GetImageUrl(),
	}
}

func MessageDomainToProto(m *chat.Message) *chatpb.Message {
	p := &chatpb.Message{
		Id:       m.ID,
		ChatId:   m.ChatID,
		SenderId: m.SenderID,
		SentAt:   timestamppb.New(m.SentAt),
	}

	// Changed m.Content to m.Payload
	switch c := m.Payload.(type) {
	case *chat.StandardMessage: // Changed chat.StandardContent to chat.StandardMessage
		p.MessagePayload = &chatpb.Message_Standard{
			Standard: &chatpb.StandardMessage{Content: c.Content},
		}
	case *chat.QuestionMessage: // Changed chat.QuestionContent to chat.QuestionMessage
		p.MessagePayload = &chatpb.Message_Question{
			Question: &chatpb.QuestionMessage{Content: c.Content, Tags: c.Tags},
		}
	case *chat.AnswerMessage: // Changed chat.AnswerContent to chat.AnswerMessage
		p.MessagePayload = &chatpb.Message_Answer{
			Answer: &chatpb.AnswerMessage{Content: c.Content, QuestionId: c.QuestionID},
		}
	case *chat.PromotionalMessage: // Changed chat.PromotionalContent to chat.PromotionalMessage
		p.MessagePayload = &chatpb.Message_Promotional{
			Promotional: &chatpb.PromotionalMessage{
				Title:     c.Title,
				Body:      c.Body,
				ActionUrl: c.ActionURL,
				ImageUrl:  c.ImageURL,
			},
		}
	}

	return p
}

func ChatDomainToProto(c *chat.Chat) *chatpb.Chat {
	if c == nil {
		return nil
	}

	protoChat := &chatpb.Chat{
		Id:                 c.ID,
		GeneralUserId:      c.GeneralUserID,
		ProfessionalUserId: c.ProfessionalUserID,
		CreatedAt:          timestamppb.New(c.CreatedAt),
		UpdatedAt:          timestamppb.New(c.UpdatedAt),
	}

	if c.LatestMessage != nil {
		protoChat.LatestMessage = MessageDomainToProto(c.LatestMessage)
	}

	return protoChat
}
