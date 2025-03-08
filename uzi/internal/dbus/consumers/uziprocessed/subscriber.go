package uziprocessed

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/WantBeASleep/med_ml_lib/dbus"
	"github.com/google/uuid"

	pb "uzi/internal/generated/dbus/consume/uziprocessed"
	"uzi/internal/services"
	"uzi/internal/services/node_segment"
)

var ErrInvalidContor = errors.New("invalid contor")

type subscriber struct {
	services *services.Services
}

func New(
	services *services.Services,
) dbus.Consumer[*pb.UziProcessed] {
	return &subscriber{
		services: services,
	}
}

func (h *subscriber) Consume(ctx context.Context, message *pb.UziProcessed) error {
	if _, err := uuid.Parse(message.UziId); err != nil {
		return fmt.Errorf("uzi id is not uuid: %s", message.UziId)
	}

	for _, v := range message.NodesWithSegments {
		for _, segment := range v.Segments {
			if _, err := uuid.Parse(segment.ImageId); err != nil {
				return fmt.Errorf("image id is not uuid: %s", segment.ImageId)
			}
		}
	}

	arg := make([]node_segment.CreateNodesWithSegmentsArg, 0, len(message.NodesWithSegments))
	for _, v := range message.NodesWithSegments {
		node := node_segment.CreateNodesWithSegmentsArgNode{
			Ai:       v.Node.Ai,
			UziID:    uuid.MustParse(message.UziId),
			Tirads23: v.Node.Tirads_23,
			Tirads4:  v.Node.Tirads_4,
			Tirads5:  v.Node.Tirads_5,
		}

		segments := make([]node_segment.CreateNodesWithSegmentsArgSegment, 0, len(v.Segments))
		for _, segment := range v.Segments {
			// contor json parse
			if !json.Valid(segment.Contor) {
				return ErrInvalidContor
			}

			segments = append(segments, node_segment.CreateNodesWithSegmentsArgSegment{
				ImageID:  uuid.MustParse(segment.ImageId),
				Contor:   json.RawMessage(segment.Contor),
				Tirads23: segment.Tirads_23,
				Tirads4:  segment.Tirads_4,
				Tirads5:  segment.Tirads_5,
			})
		}

		arg = append(arg, node_segment.CreateNodesWithSegmentsArg{
			Node:     node,
			Segments: segments,
		})
	}

	h.services.NodeSegment.CreateNodesWithSegments(ctx, arg)
	return nil
}
