package imagesegmentsnode

import (
	"context"

	pb "uzi/internal/generated/grpc/service"

	isnSrv "uzi/internal/services/image-segment-node"

	nodemapper "uzi/internal/controllers/grpc/node"
	segmentmapper "uzi/internal/controllers/grpc/segment"

	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type ImageSegmentsNodeHandler interface {
	GetImageSegmentsWithNodes(ctx context.Context, in *pb.GetImageSegmentsWithNodesIn) (*pb.GetImageSegmentsWithNodesOut, error)
}

type handler struct {
	isnSrv isnSrv.Service
}

func New(
	isnSrv isnSrv.Service,
) ImageSegmentsNodeHandler {
	return &handler{
		isnSrv: isnSrv,
	}
}

// TODO: вынести это в сегменты или ноды, однозначно не в image
func (h *handler) GetImageSegmentsWithNodes(ctx context.Context, in *pb.GetImageSegmentsWithNodesIn) (*pb.GetImageSegmentsWithNodesOut, error) {
	nodes, segments, err := h.isnSrv.GetImageSegmentsWithNodes(ctx, uuid.MustParse(in.Id))
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Что то пошло не так: %s", err.Error())
	}

	pbnodes := make([]*pb.Node, 0, len(nodes))
	for _, v := range nodes {
		pbnodes = append(pbnodes, nodemapper.DomainNodeToPb(&v))
	}

	pbsegments := make([]*pb.Segment, 0, len(segments))
	for _, v := range segments {
		pbsegments = append(pbsegments, segmentmapper.DomainSegmentToPb(&v))
	}

	return &pb.GetImageSegmentsWithNodesOut{
		Nodes:    pbnodes,
		Segments: pbsegments,
	}, nil
}
