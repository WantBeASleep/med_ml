package mappers

import (
	"time"

	"github.com/google/uuid"

	domain "gateway/internal/domain/uzi"
	pb "gateway/internal/generated/grpc/clients/uzi"
)

var uziStatusMap = map[pb.UziStatus]domain.UziStatus{
	pb.UziStatus_UZI_STATUS_NEW:       domain.UziStatusNew,
	pb.UziStatus_UZI_STATUS_PENDING:   domain.UziStatusPending,
	pb.UziStatus_UZI_STATUS_COMPLETED: domain.UziStatusCompleted,
}

func Uzi(pb *pb.Uzi) domain.Uzi {
	createAt, _ := time.Parse(time.RFC3339, pb.CreateAt)

	return domain.Uzi{
		Id:         uuid.MustParse(pb.Id),
		Projection: pb.Projection,
		Checked:    pb.Checked,
		ExternalID: uuid.MustParse(pb.ExternalId),
		DeviceID:   int(pb.DeviceId),
		Status:     uziStatusMap[pb.Status],
		CreateAt:   createAt,
	}
}

func SliceUzi(pbs []*pb.Uzi) []domain.Uzi {
	domains := make([]domain.Uzi, 0, len(pbs))
	for _, pb := range pbs {
		domains = append(domains, Uzi(pb))
	}
	return domains
}
