package grpc

import (
	"uzi/internal/generated/grpc/service"
	"uzi/internal/server/device"
	"uzi/internal/server/image"
	"uzi/internal/server/node"
	"uzi/internal/server/node_segment"
	"uzi/internal/server/segment"
	"uzi/internal/server/uzi"
	"uzi/internal/services"
)

// из за эмбедина приходится делать приписку перед Handler
type Handler struct {
	device.DeviceHandler
	uzi.UziHandler
	image.ImageHandler
	node_segment.NodeSegmentHandler
	node.NodeHandler
	segment.SegmentHandler

	service.UnsafeUziSrvServer
}

func New(
	services *services.Services,
) *Handler {
	deviceHandler := device.New(services)
	uziHandler := uzi.New(services)
	imageHandler := image.New(services)
	nodeSegmentHandler := node_segment.New(services)
	nodeHandler := node.New(services)
	segmentHandler := segment.New(services)

	return &Handler{
		DeviceHandler:      deviceHandler,
		UziHandler:         uziHandler,
		ImageHandler:       imageHandler,
		NodeSegmentHandler: nodeSegmentHandler,
		NodeHandler:        nodeHandler,
		SegmentHandler:     segmentHandler,
	}
}
