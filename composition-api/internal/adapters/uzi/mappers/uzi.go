package mappers

import (
	"time"

	"github.com/google/uuid"

	domain "composition-api/internal/domain/uzi"
	pb "composition-api/internal/generated/grpc/clients/uzi"
)

var uziStatusMap = map[pb.UziStatus]domain.UziStatus{
	pb.UziStatus_UZI_STATUS_NEW:       domain.UziStatusNew,
	pb.UziStatus_UZI_STATUS_PENDING:   domain.UziStatusPending,
	pb.UziStatus_UZI_STATUS_COMPLETED: domain.UziStatusCompleted,
}

var uziProjectionMap = map[pb.UziProjection]domain.UziProjection{
	pb.UziProjection_UZI_PROJECTION_CROSS: domain.UziProjectionCross,
	pb.UziProjection_UZI_PROJECTION_LONG:  domain.UziProjectionLong,
}

type Uzi struct{}

func (m Uzi) Domain(pb *pb.Uzi) domain.Uzi {
	createAt, _ := time.Parse(time.RFC3339, pb.CreateAt)

	return domain.Uzi{
		Id:          uuid.MustParse(pb.Id),
		Projection:  uziProjectionMap[pb.Projection],
		Checked:     pb.Checked,
		ExternalID:  uuid.MustParse(pb.ExternalId),
		Author:      uuid.MustParse(pb.Author),
		DeviceID:    int(pb.DeviceId),
		Status:      uziStatusMap[pb.Status],
		Description: pb.Description,
		CreateAt:    createAt,
	}
}

func (m Uzi) SliceDomain(pbs []*pb.Uzi) []domain.Uzi {
	return slice(pbs, m)
}
