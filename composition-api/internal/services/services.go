package services

import (
	"composition-api/internal/services/device"
	"composition-api/internal/services/image"
	"composition-api/internal/services/node"
	"composition-api/internal/services/node_segment"
	"composition-api/internal/services/segment"
	"composition-api/internal/services/uzi"

	"composition-api/internal/adapters"
	"composition-api/internal/dbus/producers"
	"composition-api/internal/repository"
)

type Services struct {
	DeviceService      device.Service
	UziService         uzi.Service
	ImageService       image.Service
	NodeService        node.Service
	SegmentService     segment.Service
	NodeSegmentService node_segment.Service
}

func New(
	adapters *adapters.Adapters,
	producers producers.Producer,
	dao repository.DAO,
) *Services {
	deviceService := device.New(adapters)
	uziService := uzi.New(adapters, dao, producers)
	imageService := image.New(adapters)
	nodeService := node.New(adapters)
	segmentService := segment.New(adapters)
	nodeSegmentService := node_segment.New(adapters)

	return &Services{
		DeviceService:      deviceService,
		UziService:         uziService,
		ImageService:       imageService,
		NodeService:        nodeService,
		SegmentService:     segmentService,
		NodeSegmentService: nodeSegmentService,
	}
}
