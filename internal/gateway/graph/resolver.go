//go:generate go run github.com/99designs/gqlgen generate
package graph

import (
	chatpb "github.com/maehiyu/tollo/gen/go/protos/chatservice"
	userpb "github.com/maehiyu/tollo/gen/go/protos/userservice"
	"github.com/maehiyu/tollo/internal/gateway/graph/model" // Added for model types
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	UserClient userpb.UserServiceClient
	ChatClient chatpb.ChatServiceClient
}

// Helper method to convert protobuf Chat to GraphQL Chat
func (r *Resolver) ProtoChatToGraphQLChat(protoChat *chatpb.Chat) *model.Chat {
	if protoChat == nil {
		return nil
	}

	gqlChat := &model.Chat{
		ID:                 protoChat.Id,
		GeneralUserID:      protoChat.GeneralUserId,
		ProfessionalUserID: protoChat.ProfessionalUserId,
		CreatedAt:          protoChat.CreatedAt.AsTime(),
		UpdatedAt:          protoChat.UpdatedAt.AsTime(),
	}

	if protoChat.LatestMessage != nil {
		gqlChat.LatestMessage = r.ProtoMessageToGraphQLMessage(protoChat.LatestMessage)
	}

	return gqlChat
}

// Helper method to convert protobuf Message to GraphQL Message
func (r *Resolver) ProtoMessageToGraphQLMessage(protoMsg *chatpb.Message) *model.Message {
	if protoMsg == nil {
		return nil
	}

	gqlMsg := &model.Message{
		ID:       protoMsg.Id,
		ChatID:   protoMsg.ChatId,
		SenderID: protoMsg.SenderId,
		SentAt:   protoMsg.SentAt.AsTime(),
	}

	// Handle message_payload oneof
	switch payload := protoMsg.MessagePayload.(type) {
	case *chatpb.Message_Standard:
		gqlMsg.Payload = &model.StandardMessage{Content: payload.Standard.Content}
	case *chatpb.Message_Question:
		gqlMsg.Payload = &model.QuestionMessage{Content: payload.Question.Content, Tags: payload.Question.Tags}
	case *chatpb.Message_Answer:
		gqlMsg.Payload = &model.AnswerMessage{Content: payload.Answer.Content, QuestionID: payload.Answer.QuestionId}
	case *chatpb.Message_Promotional:
		var imageUrl *string
		if payload.Promotional.ImageUrl != "" {
			imageUrl = &payload.Promotional.ImageUrl
		}
		gqlMsg.Payload = &model.PromotionalMessage{
			Title:     payload.Promotional.Title,
			Body:      payload.Promotional.Body,
			ActionURL: payload.Promotional.ActionUrl,
			ImageURL:  imageUrl,
		}
	}

	return gqlMsg
}

// Helper method to convert protobuf UserInfo to GraphQL User
func (r *Resolver) ProtoUserToGraphQLUser(protoUser *userpb.UserInfo) *model.User {
	if protoUser == nil {
		return nil
	}

	gqlUser := &model.User{
		ID:        protoUser.Id,
		Name:      protoUser.Name,
		Email:     protoUser.Email,
		CreatedAt: protoUser.CreatedAt.AsTime(),
		UpdatedAt: protoUser.UpdatedAt.AsTime(),
	}

	// Handle the profile oneof from the response
	if profProfile := protoUser.GetProfessionalProfile(); profProfile != nil {
		gqlUser.Profile = &model.ProfessionalProfile{
			ProBadgeURL: profProfile.ProBadgeUrl,
			Biography:   profProfile.Biography,
		}
	} else if genProfile := protoUser.GetGeneralProfile(); genProfile != nil {
		gqlUser.Profile = &model.GeneralProfile{
			Points:       int32(genProfile.Points),
			Introduction: genProfile.Introduction,
		}
	}

	return gqlUser
}
