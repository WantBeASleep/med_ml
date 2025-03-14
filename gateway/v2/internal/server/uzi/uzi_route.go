package uzi

import (
	"gateway/internal/server/uzi/device"
	"gateway/internal/server/uzi/image"
	"gateway/internal/server/uzi/node"
	"gateway/internal/server/uzi/node_segment"
	"gateway/internal/server/uzi/segment"
	"gateway/internal/server/uzi/uzi"
	services "gateway/internal/services"
)

type UziRoute interface {
	segment.SegmentHandler
	node_segment.NodeSegmentHandler
	device.DeviceHandler
	image.ImageHandler
	node.NodeHandler
	uzi.UziHandler
}

type uziRoute struct {
	segment.SegmentHandler
	node_segment.NodeSegmentHandler
	device.DeviceHandler
	image.ImageHandler
	node.NodeHandler
	uzi.UziHandler
}

func NewUziRoute(services *services.Services) UziRoute {
	segmentHandler := segment.NewHandler(services)
	nodeSegmentHandler := node_segment.NewHandler(services)
	deviceHandler := device.NewHandler(services)
	imageHandler := image.NewHandler(services)
	nodeHandler := node.NewHandler(services)
	uziHandler := uzi.NewHandler(services)

	return &uziRoute{
		SegmentHandler:     segmentHandler,
		NodeSegmentHandler: nodeSegmentHandler,
		DeviceHandler:      deviceHandler,
		ImageHandler:       imageHandler,
		NodeHandler:        nodeHandler,
		UziHandler:         uziHandler,
	}
}
