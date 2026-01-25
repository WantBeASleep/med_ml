package mappers

import (
	"time"

	"github.com/google/uuid"

	domain "composition-api/internal/domain/cytology"
	pb "composition-api/internal/generated/grpc/clients/cytology"
)

var segTypeMap = map[pb.SegType]domain.SegType{
	pb.SegType_SEG_TYPE_NIL: domain.SegTypeNIL,
	pb.SegType_SEG_TYPE_NIR: domain.SegTypeNIR,
	pb.SegType_SEG_TYPE_NIM: domain.SegTypeNIM,
	pb.SegType_SEG_TYPE_CNO: domain.SegTypeCNO,
	pb.SegType_SEG_TYPE_CGE: domain.SegTypeCGE,
	pb.SegType_SEG_TYPE_C2N: domain.SegTypeC2N,
	pb.SegType_SEG_TYPE_CPS: domain.SegTypeCPS,
	pb.SegType_SEG_TYPE_CFC: domain.SegTypeCFC,
	pb.SegType_SEG_TYPE_CLY: domain.SegTypeCLY,
	pb.SegType_SEG_TYPE_SOS: domain.SegTypeSOS,
	pb.SegType_SEG_TYPE_SDS: domain.SegTypeSDS,
	pb.SegType_SEG_TYPE_SMS: domain.SegTypeSMS,
	pb.SegType_SEG_TYPE_STS: domain.SegTypeSTS,
	pb.SegType_SEG_TYPE_SPS: domain.SegTypeSPS,
	pb.SegType_SEG_TYPE_SNM: domain.SegTypeSNM,
	pb.SegType_SEG_TYPE_STM: domain.SegTypeSTM,
}

var groupTypeMap = map[pb.GroupType]domain.GroupType{
	pb.GroupType_GROUP_TYPE_CE: domain.GroupTypeCE,
	pb.GroupType_GROUP_TYPE_CL: domain.GroupTypeCL,
	pb.GroupType_GROUP_TYPE_ME: domain.GroupTypeME,
}

type SegmentationGroup struct{}

func (m SegmentationGroup) Domain(pb *pb.SegmentationGroup) domain.SegmentationGroup {
	createAt, _ := time.Parse(time.RFC3339, pb.CreateAt)

	return domain.SegmentationGroup{
		Id:         int(pb.Id),
		CytologyID: uuid.MustParse(pb.CytologyId),
		SegType:    segTypeMap[pb.SegType],
		GroupType:  groupTypeMap[pb.GroupType],
		IsAI:       pb.IsAi,
		Details:    pb.Details,
		CreateAt:   createAt,
	}
}

func (m SegmentationGroup) SliceDomain(pbs []*pb.SegmentationGroup) []domain.SegmentationGroup {
	domains := make([]domain.SegmentationGroup, 0, len(pbs))
	for _, pb := range pbs {
		domains = append(domains, m.Domain(pb))
	}
	return domains
}

type Segmentation struct{}

func (m Segmentation) Domain(pb *pb.Segmentation) domain.Segmentation {
	createAt, _ := time.Parse(time.RFC3339, pb.CreateAt)

	points := make([]domain.SegmentationPoint, 0, len(pb.Points))
	for _, p := range pb.Points {
		pointCreateAt, _ := time.Parse(time.RFC3339, "")
		points = append(points, domain.SegmentationPoint{
			Id:             int(p.Id),
			SegmentationID: int(p.SegmentationId),
			X:              int(p.X),
			Y:              int(p.Y),
			UID:            p.Uid,
			CreateAt:       pointCreateAt,
		})
	}

	return domain.Segmentation{
		Id:                  int(pb.Id),
		SegmentationGroupID: int(pb.SegmentationGroupId),
		Points:              points,
		CreateAt:            createAt,
	}
}

func (m Segmentation) SliceDomain(pbs []*pb.Segmentation) []domain.Segmentation {
	domains := make([]domain.Segmentation, 0, len(pbs))
	for _, pb := range pbs {
		domains = append(domains, m.Domain(pb))
	}
	return domains
}
