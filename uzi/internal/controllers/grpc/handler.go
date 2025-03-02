package grpc

import (
	"uzi/internal/controllers/grpc/device"
	"uzi/internal/controllers/grpc/image"
	isn "uzi/internal/controllers/grpc/image-segment-node"
	"uzi/internal/controllers/grpc/node"
	"uzi/internal/controllers/grpc/segment"
	"uzi/internal/controllers/grpc/uzi"
	"uzi/internal/generated/grpc/service"
)

type Handler struct {
	device.DeviceHandler
	uzi.UziHandler
	image.ImageHandler
	isn.ImageSegmentsNodeHandler
	node.NodeHandler
	segment.SegmentHandler

	service.UnsafeUziSrvServer
}

func New(
	deviceHandler device.DeviceHandler,
	uziHandler uzi.UziHandler,
	imageHandler image.ImageHandler,
	isnHandler isn.ImageSegmentsNodeHandler,
	segmentHandler segment.SegmentHandler,
	nodeHandler node.NodeHandler,
) *Handler {
	return &Handler{
		DeviceHandler:            deviceHandler,
		UziHandler:               uziHandler,
		ImageHandler:             imageHandler,
		ImageSegmentsNodeHandler: isnHandler,
		NodeHandler:              nodeHandler,
		SegmentHandler:           segmentHandler,
	}
}
