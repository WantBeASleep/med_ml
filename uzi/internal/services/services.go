package services

import (
	"uzi/internal/repository"
	"uzi/internal/services/device"
	"uzi/internal/services/image"
	"uzi/internal/services/node"
	"uzi/internal/services/node_segment"
	"uzi/internal/services/segment"
	"uzi/internal/services/splitter"
	"uzi/internal/services/uzi"

	dbus "uzi/internal/dbus/producers"
)

type Services struct {
	Device      device.Service
	Uzi         uzi.Service
	Image       image.Service
	Node        node.Service
	Segment     segment.Service
	NodeSegment node_segment.Service
	Splitter    splitter.Service
}

func New(
	dao repository.DAO,
	dbus dbus.Producer,
) *Services {
	device := device.New(dao)
	uzi := uzi.New(dao)
	image := image.New(dao, dbus)
	node := node.New(dao)
	segment := segment.New(dao)
	nodeSegment := node_segment.New(dao)
	splitter := splitter.New()

	return &Services{
		Device:      device,
		Uzi:         uzi,
		Image:       image,
		Node:        node,
		Segment:     segment,
		NodeSegment: nodeSegment,
		Splitter:    splitter,
	}
}
