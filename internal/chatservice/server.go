package chatservice

import (
	"context"

	"github.com/maehiyu/tollo/internal/adapter/converter"
	"github.com/maehiyu/tollo/internal/chatservice/domain/chat"
	pb "github.com/maehiyu/tollo/gen/go/protos/chatservice"
)

var _ pb.ChatServiceServer = (*Server)(nil)

type Server struct {
	pb.UnimplementedChatServiceServer
	usecase *Usecase
}

func NewServer(usecase *Usecase) *Server {
	return &Server{
		usecase: usecase,
	}
}

func (s *Server) SendMessage(ctx context.Context, req *pb.SendMessageRequest) (*pb.SendMessageResponse, error) {
	var payload chat.MessagePayload

	switch p := req.MessagePayload.(type) {
	case *pb.SendMessageRequest_Standard:
		payload = converter.StandardPayloadProtoToDomain(p.Standard)
	case *pb.SendMessageRequest_Question:
		payload = converter.QuestionPayloadProtoToDomain(p.Question)
	case *pb.SendMessageRequest_Answer:
		payload = converter.AnswerPayloadProtoToDomain(p.Answer)
	case *pb.SendMessageRequest_Promotional:
		payload = converter.PromotionalPayloadProtoToDomain(p.Promotional)
	}

	sentMessage, err := s.usecase.SendMessage(ctx, req.GetChatId(), req.GetSenderId(), payload)
	if err != nil {
		return nil, err
	}

	return &pb.SendMessageResponse{
		SentMessage: converter.MessageDomainToProto(sentMessage),
	}, nil
}

func (s *Server) GetUserChats(ctx context.Context, req *pb.GetUserChatsRequest) (*pb.GetUserChatsResponse, error) {
	userID := req.GetUserId()

	domainChats, err := s.usecase.GetChatsByUserID(ctx, userID)
	if err != nil {
		return nil, err
	}

	protoChats := make([]*pb.Chat, len(domainChats))
	for i, c := range domainChats {
		protoChats[i] = converter.ChatDomainToProto(c)
	}

	return &pb.GetUserChatsResponse{
		Chats: protoChats,
	}, nil
}

func (s *Server) GetChatMessages(ctx context.Context, req *pb.GetChatMessagesRequest) (*pb.GetChatMessagesResponse, error) {
	chatID := req.GetChatId()

	domainMessages, err := s.usecase.GetMessagesByChatID(ctx, chatID)
	if err != nil {
		return nil, err
	}

	protoMessages := make([]*pb.Message, len(domainMessages))
	for i, m := range domainMessages {
		protoMessages[i] = converter.MessageDomainToProto(m)
	}

	return &pb.GetChatMessagesResponse{
		Messages: protoMessages,
	}, nil
}

func (s *Server) CreateChat(ctx context.Context, req *pb.CreateChatRequest) (*pb.CreateChatResponse, error) {
	generalUserID := req.GetGeneralUserId()
	professionalUserID := req.GetProfessionalUserId()

	domainChat, err := s.usecase.CreateChat(ctx, generalUserID, professionalUserID)
	if err != nil {
		return nil, err
	}

	return &pb.CreateChatResponse{
		Chat: converter.ChatDomainToProto(domainChat),
	}, nil
}
