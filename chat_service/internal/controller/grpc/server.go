package grpc

import (
	"fmt"
	"net"

	chatv1 "chat_service/internal/pb/chat/v1"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type Server struct {
	server      *grpc.Server
	address     string
	chatUsecase ChatUsecase
	chatv1.UnimplementedChatSrvServer
}

func NewServer(address string, chatUsecase ChatUsecase) (*Server, error) {
	s := &Server{
		server:      grpc.NewServer(),
		address:     address,
		chatUsecase: chatUsecase,
	}

	chatv1.RegisterChatSrvServer(s.server, s)

	reflection.Register(s.server)

	return s, nil
}

func (s *Server) Start() error {
	listener, err := net.Listen("tcp", s.address)
	if err != nil {
		return fmt.Errorf("listen: %w", err)
	}

	return s.server.Serve(listener)
}

func (s *Server) Stop() {
	s.server.GracefulStop()
}
