package mappers

import (
	"time"

	"github.com/google/uuid"

	domain "composition-api/internal/domain/cytology"
	pb "composition-api/internal/generated/grpc/clients/cytology"
)

type OriginalImage struct{}

func (m OriginalImage) Domain(pb *pb.OriginalImage) domain.OriginalImage {
	createDate, _ := time.Parse(time.RFC3339, pb.CreateDate)

	var delayTime *float64
	if pb.DelayTime != nil {
		dt := *pb.DelayTime
		delayTime = &dt
	}

	return domain.OriginalImage{
		Id:         uuid.MustParse(pb.Id),
		CytologyID: uuid.MustParse(pb.CytologyId),
		ImagePath:  pb.ImagePath,
		CreateDate: createDate,
		DelayTime:  delayTime,
		ViewedFlag: pb.ViewedFlag,
	}
}

func (m OriginalImage) SliceDomain(pbs []*pb.OriginalImage) []domain.OriginalImage {
	domains := make([]domain.OriginalImage, 0, len(pbs))
	for _, pb := range pbs {
		domains = append(domains, m.Domain(pb))
	}
	return domains
}
