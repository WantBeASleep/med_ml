package mappers

import (
	"cytology/internal/domain"
	pb "cytology/internal/generated/grpc/service"
)

var SegTypeMap = map[domain.SegType]pb.SegType{
	domain.SegTypeNIL: pb.SegType_SEG_TYPE_NIL,
	domain.SegTypeNIR: pb.SegType_SEG_TYPE_NIR,
	domain.SegTypeNIM: pb.SegType_SEG_TYPE_NIM,
	domain.SegTypeCNO: pb.SegType_SEG_TYPE_CNO,
	domain.SegTypeCGE: pb.SegType_SEG_TYPE_CGE,
	domain.SegTypeC2N: pb.SegType_SEG_TYPE_C2N,
	domain.SegTypeCPS: pb.SegType_SEG_TYPE_CPS,
	domain.SegTypeCFC: pb.SegType_SEG_TYPE_CFC,
	domain.SegTypeCLY: pb.SegType_SEG_TYPE_CLY,
	domain.SegTypeSOS: pb.SegType_SEG_TYPE_SOS,
	domain.SegTypeSDS: pb.SegType_SEG_TYPE_SDS,
	domain.SegTypeSMS: pb.SegType_SEG_TYPE_SMS,
	domain.SegTypeSTS: pb.SegType_SEG_TYPE_STS,
	domain.SegTypeSPS: pb.SegType_SEG_TYPE_SPS,
	domain.SegTypeSNM: pb.SegType_SEG_TYPE_SNM,
	domain.SegTypeSTM: pb.SegType_SEG_TYPE_STM,
}

var SegTypeReverseMap = map[pb.SegType]domain.SegType{
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

var GroupTypeMap = map[domain.GroupType]pb.GroupType{
	domain.GroupTypeCE: pb.GroupType_GROUP_TYPE_CE,
	domain.GroupTypeCL: pb.GroupType_GROUP_TYPE_CL,
	domain.GroupTypeME: pb.GroupType_GROUP_TYPE_ME,
}

var GroupTypeReverseMap = map[pb.GroupType]domain.GroupType{
	pb.GroupType_GROUP_TYPE_CE: domain.GroupTypeCE,
	pb.GroupType_GROUP_TYPE_CL: domain.GroupTypeCL,
	pb.GroupType_GROUP_TYPE_ME: domain.GroupTypeME,
}

func SegmentationGroupToProto(d domain.SegmentationGroup) *pb.SegmentationGroup {
	pbGroup := &pb.SegmentationGroup{
		Id:         int32(d.Id),
		CytologyId: d.CytologyID.String(),
		SegType:    SegTypeMap[d.SegType],
		GroupType:  GroupTypeMap[d.GroupType],
		IsAi:       d.IsAI,
		CreateAt:   d.CreateAt.Format("2006-01-02T15:04:05Z07:00"),
	}

	if d.Details != nil && len(d.Details) > 0 {
		details := string(d.Details)
		pbGroup.Details = &details
	}

	return pbGroup
}

func SegmentationGroupSliceToProto(groups []domain.SegmentationGroup) []*pb.SegmentationGroup {
	pbGroups := make([]*pb.SegmentationGroup, 0, len(groups))
	for _, group := range groups {
		pbGroups = append(pbGroups, SegmentationGroupToProto(group))
	}
	return pbGroups
}

func SegmentationToProto(d domain.Segmentation) *pb.Segmentation {
	pbSeg := &pb.Segmentation{
		Id:                  int32(d.Id),
		SegmentationGroupId: int32(d.SegmentationGroupID),
		CreateAt:            d.CreateAt.Format("2006-01-02T15:04:05Z07:00"),
		Points:              make([]*pb.SegmentationPoint, 0, len(d.Points)),
	}

	for _, point := range d.Points {
		pbSeg.Points = append(pbSeg.Points, &pb.SegmentationPoint{
			Id:             int32(point.Id),
			SegmentationId: int32(point.SegmentationID),
			X:              int32(point.X),
			Y:              int32(point.Y),
			Uid:            point.UID,
		})
	}

	return pbSeg
}

func SegmentationSliceToProto(segs []domain.Segmentation) []*pb.Segmentation {
	pbSegs := make([]*pb.Segmentation, 0, len(segs))
	for _, seg := range segs {
		pbSegs = append(pbSegs, SegmentationToProto(seg))
	}
	return pbSegs
}
