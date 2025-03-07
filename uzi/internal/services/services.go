package services

import (
	"uzi/internal/repository"
	"uzi/internal/services/device"
	"uzi/internal/services/image"
	"uzi/internal/services/image_segment_node"
	"uzi/internal/services/node"
	"uzi/internal/services/segment"
	"uzi/internal/services/splitter"
	"uzi/internal/services/uzi"

	"uzi/internal/adapters/dbus"
)

type Services struct {
	Device           device.Service
	Uzi              uzi.Service
	Image            image.Service
	Node             node.Service
	Segment          segment.Service
	ImageSegmentNode image_segment_node.Service
	Splitter         splitter.Service
}

func NewServices(
	dao repository.DAO,
	dbusAdapter dbus.DbusAdapter,
) *Services {
	device := device.New(dao)
	uzi := uzi.New(dao)
	image := image.New(dao, dbusAdapter)
	node := node.New(dao)
	segment := segment.New(dao)
	imageSegmentNode := image_segment_node.New(dao)
	splitter := splitter.New()

	return &Services{
		Device:           device,
		Uzi:              uzi,
		Image:            image,
		Node:             node,
		Segment:          segment,
		ImageSegmentNode: imageSegmentNode,
		Splitter:         splitter,
	}
}
