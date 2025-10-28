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

func StandardMessageProtoToDomain(p *chatpb.StandardMessage) *chat.StandardContent {
	return &chat.StandardContent{
		Content: p.GetContent(),
	}
}

func QuestionMessageProtoToDomain(p *chatpb.QuestionMessage) *chat.QuestionContent {
	return &chat.QuestionContent{
		Content: p.GetContent(),
		Tags:    p.GetTags(),
	}
}

func AnswerMessageProtoToDomain(p *chatpb.AnswerMessage) *chat.AnswerContent {
	return &chat.AnswerContent{
		Content:    p.GetContent(),
		QuestionID: p.GetQuestionId(),
	}
}

func PromotionalMessageProtoToDomain(p *chatpb.PromotionalMessage) *chat.PromotionalContent {
	return &chat.PromotionalContent{
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

	switch c := m.Content.(type) {
	case *chat.StandardContent:
		p.MessagePayload = &chatpb.Message_Standard{
			Standard: &chatpb.StandardMessage{Content: c.Content},
		}
	case *chat.QuestionContent:
		p.MessagePayload = &chatpb.Message_Question{
			Question: &chatpb.QuestionMessage{Content: c.Content, Tags: c.Tags},
		}
	case *chat.AnswerContent:
		p.MessagePayload = &chatpb.Message_Answer{
			Answer: &chatpb.AnswerMessage{Content: c.Content, QuestionId: c.QuestionID},
		}
	case *chat.PromotionalContent:
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
