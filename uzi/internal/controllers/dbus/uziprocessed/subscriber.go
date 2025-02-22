package uziprocessed

import (
	"context"
	"fmt"

	"github.com/WantBeASleep/med_ml_lib/dbus"

	"uzi/internal/domain"
	pb "uzi/internal/generated/dbus/consume/uziprocessed"
	"uzi/internal/services/node"

	"github.com/google/uuid"
)

type subscriber struct {
	nodeSrv node.Service
}

func New(
	nodeSrv node.Service,
) dbus.Consumer[*pb.UziProcessed] {
	return &subscriber{
		nodeSrv: nodeSrv,
	}
}

func (h *subscriber) Consume(ctx context.Context, event *pb.UziProcessed) error {
	nodes := make([]domain.Node, 0, len(event.Nodes))
	segments := make([]domain.Segment, 0, len(event.Segments))

	for _, v := range event.Nodes {
		nodes = append(nodes, domain.Node{
			Id:       uuid.MustParse(v.Id),
			UziID:    uuid.MustParse(v.UziId),
			Tirads23: v.Tirads_23,
			Tirads4:  v.Tirads_4,
			Tirads5:  v.Tirads_5,
		})
	}

	for _, v := range event.Segments {
		segments = append(segments, domain.Segment{
			Id:       uuid.MustParse(v.Id),
			ImageID:  uuid.MustParse(v.ImageId),
			NodeID:   uuid.MustParse(v.NodeId),
			Contor:   v.Contor,
			Tirads23: v.Tirads_23,
			Tirads4:  v.Tirads_4,
			Tirads5:  v.Tirads_5,
		})
	}

	if err := h.nodeSrv.InsertAiNodeWithSegments(ctx, nodes, segments); err != nil {
		return fmt.Errorf("isert ai nodes && segments: %w", err)
	}
	return nil
}
