package services

import (
	"gateway/internal/services/device"
	"gateway/internal/services/image"
	"gateway/internal/services/node"
	"gateway/internal/services/node_segment"
	"gateway/internal/services/segment"
	"gateway/internal/services/uzi"

	"gateway/internal/adapters"
	"gateway/internal/dbus/producers"
	"gateway/internal/repository"
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
