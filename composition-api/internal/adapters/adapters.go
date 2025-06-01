package adapters

import (
	"google.golang.org/grpc"

	"composition-api/internal/adapters/auth"
	"composition-api/internal/adapters/chat"
	"composition-api/internal/adapters/med"
	"composition-api/internal/adapters/uzi"
	authPB "composition-api/internal/generated/grpc/clients/auth"
	chatPB "composition-api/internal/generated/grpc/clients/chat"
	medPB "composition-api/internal/generated/grpc/clients/med"
	uziPB "composition-api/internal/generated/grpc/clients/uzi"
)

type Adapters struct {
	Uzi  uzi.Adapter
	Auth auth.Adapter
	Med  med.Adapter
	Chat chat.Adapter
}

func NewAdapters(
	uziConn *grpc.ClientConn,
	authConn *grpc.ClientConn,
	medConn *grpc.ClientConn,
	chatConn *grpc.ClientConn,
) *Adapters {
	uziClient := uziPB.NewUziSrvClient(uziConn)
	uziAdapter := uzi.NewAdapter(uziClient)

	authClient := authPB.NewAuthSrvClient(authConn)
	authAdapter := auth.NewAdapter(authClient)

	medClient := medPB.NewMedSrvClient(medConn)
	medAdapter := med.NewAdapter(medClient)

	chatClient := chatPB.NewChatSrvClient(chatConn)
	chatAdapter := chat.NewAdapter(chatClient)

	return &Adapters{
		Uzi:  uziAdapter,
		Auth: authAdapter,
		Med:  medAdapter,
		Chat: chatAdapter,
	}
}
