package segment

import (
	"context"
	"encoding/json"

	"github.com/AlekSi/pointer"
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	pb "uzi/internal/generated/grpc/service"
	"uzi/internal/server/mappers"
	"uzi/internal/services/segment"
)

func (h *handler) UpdateSegment(ctx context.Context, in *pb.UpdateSegmentIn) (*pb.UpdateSegmentOut, error) {
	if _, err := uuid.Parse(in.Id); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "id is not a valid uuid: %s", err.Error())
	}

	var contor *json.RawMessage
	if in.Contor != nil {
		if !json.Valid(in.Contor) {
			return nil, status.Errorf(codes.InvalidArgument, "contor is not a valid json")
		}
		contor = pointer.To(json.RawMessage(in.Contor))
	}

	segment, err := h.services.Segment.UpdateSegment(
		ctx,
		segment.UpdateSegmentArg{
			Id:       uuid.MustParse(in.Id),
			Contor:   contor,
			Tirads23: in.Tirads_23,
			Tirads4:  in.Tirads_4,
			Tirads5:  in.Tirads_5,
		},
	)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Что то пошло не так: %s", err.Error())
	}

	out := new(pb.UpdateSegmentOut)
	out.Segment = mappers.SegmentFromDomain(segment)

	return out, nil
}
