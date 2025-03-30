package mappers

import (
	"time"

	"uzi/internal/domain"
	pb "uzi/internal/generated/grpc/service"
)

var uziStatusMap = map[domain.UziStatus]pb.UziStatus{
	domain.UziStatusNew:       pb.UziStatus_UZI_STATUS_NEW,
	domain.UziStatusPending:   pb.UziStatus_UZI_STATUS_PENDING,
	domain.UziStatusCompleted: pb.UziStatus_UZI_STATUS_COMPLETED,
}

var UziProjectionMap = map[domain.UziProjection]pb.UziProjection{
	domain.UziProjectionLong:  pb.UziProjection_UZI_PROJECTION_LONG,
	domain.UziProjectionCross: pb.UziProjection_UZI_PROJECTION_CROSS,
}

var UziProjectionReverseMap = map[pb.UziProjection]domain.UziProjection{
	pb.UziProjection_UZI_PROJECTION_LONG:  domain.UziProjectionLong,
	pb.UziProjection_UZI_PROJECTION_CROSS: domain.UziProjectionCross,
}

func UziFromDomain(domain domain.Uzi) *pb.Uzi {
	return &pb.Uzi{
		Id:          domain.Id.String(),
		Projection:  UziProjectionMap[domain.Projection],
		Checked:     domain.Checked,
		ExternalId:  domain.ExternalID.String(),
		Author:      domain.Author.String(),
		DeviceId:    int64(domain.DeviceID),
		Status:      uziStatusMap[domain.Status],
		Description: domain.Description,
		CreateAt:    domain.CreateAt.Format(time.RFC3339),
	}
}

func SliceUziFromDomain(domains []domain.Uzi) []*pb.Uzi {
	pbs := make([]*pb.Uzi, 0, len(domains))
	for _, d := range domains {
		pbs = append(pbs, UziFromDomain(d))
	}
	return pbs
}
