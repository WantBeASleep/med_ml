package uzi

import (
	"composition-api/internal/server/uzi/device"
	"composition-api/internal/server/uzi/image"
	"composition-api/internal/server/uzi/node"
	"composition-api/internal/server/uzi/node_segment"
	"composition-api/internal/server/uzi/segment"
	"composition-api/internal/server/uzi/uzi"
	services "composition-api/internal/services"
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
