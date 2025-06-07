package grpc

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/WantBeASleep/med_ml_lib/observer/consts"
	"github.com/WantBeASleep/med_ml_lib/observer/cross"
	loglib "github.com/WantBeASleep/med_ml_lib/observer/log"
	"github.com/google/uuid"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

func CrossServerCall(
	ctx context.Context,
	req any,
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler,
) (resp any, err error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		slog.WarnContext(ctx, "Server call w/o metadata", slog.String("called method", info.FullMethod))
		return handler(ctx, req)
	}

	crossValues := cross.GetContextAttrs(ctx)
	for _, v := range consts.Consts {
		crossValues[v] = md.Get(v)
	}

	ctx = cross.WithFields(ctx, crossValues)

	return handler(ctx, req)
}

func CrossClientCall(
	ctx context.Context,
	method string,
	req, reply any,
	cc *grpc.ClientConn,
	invoker grpc.UnaryInvoker,
	opts ...grpc.CallOption,
) error {
	crossValues := cross.GetContextAttrs(ctx)
	newmd := map[string]string{}
	for k, v := range crossValues {
		newmd[k] = fmt.Sprintf("%v", v)
	}
	md := metadata.New(newmd)

	ctx = metadata.NewOutgoingContext(ctx, md)

	return invoker(ctx, method, req, reply, cc, opts...)
}

func LogServerCall(
	ctx context.Context,
	req any,
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler,
) (resp any, err error) {
	crossValues := cross.GetContextAttrs(ctx)

	requestIDField, ok := crossValues[consts.RequestID].([]string)
	if !ok {
		slog.WarnContext(ctx, "Server call w/o metadata", slog.String("called method", info.FullMethod))
		return handler(ctx, req)
	}
	if len(requestIDField) != 1 {
		slog.WarnContext(ctx, "Server call with wrong metadata. RequestID must be 1 field", slog.String("called method", info.FullMethod), slog.Any("requestIDField", requestIDField))
		return handler(ctx, req)
	}

	requestID, err := uuid.Parse(requestIDField[0])
	if err != nil {
		slog.WarnContext(ctx, "Server call with wrong metadata. RequestID must be uuid", slog.String("called method", info.FullMethod), slog.Any("requestIDField", requestIDField))
		return handler(ctx, req)
	}

	ctx = loglib.WithFields(ctx, map[string]any{
		consts.RequestID:     requestID,
		consts.RequestMethod: info.FullMethod,
	})
	slog.InfoContext(ctx, "Server call")

	return handler(ctx, req)
}

func LogClientCall(
	ctx context.Context,
	method string,
	req, reply any,
	cc *grpc.ClientConn,
	invoker grpc.UnaryInvoker,
	opts ...grpc.CallOption,
) error {
	slog.InfoContext(ctx, "Client call", slog.String("client method call", method))
	return invoker(ctx, method, req, reply, cc, opts...)
}
