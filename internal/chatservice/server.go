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
	var content chat.MessageContent

	switch payload := req.MessagePayload.(type) {
	case *pb.SendMessageRequest_Standard:
		content = converter.StandardMessageProtoToDomain(payload.Standard)
	case *pb.SendMessageRequest_Question:
		content = converter.QuestionMessageProtoToDomain(payload.Question)
	case *pb.SendMessageRequest_Answer:
		content = converter.AnswerMessageProtoToDomain(payload.Answer)
	case *pb.SendMessageRequest_Promotional:
		content = converter.PromotionalMessageProtoToDomain(payload.Promotional)
	}

	sentMessage, err := s.usecase.SendMessage(ctx, req.GetChatId(), req.GetSenderId(), content)
	if err != nil {
		return nil, err
	}

	return &pb.SendMessageResponse{
		SentMessage: converter.MessageDomainToProto(sentMessage),
	}, nil
}
