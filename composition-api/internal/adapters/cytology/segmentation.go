package cytology

import (
	"context"

	"composition-api/internal/adapters/cytology/mappers"
	adapter_errors "composition-api/internal/adapters/errors"
	domain "composition-api/internal/domain/cytology"
	pb "composition-api/internal/generated/grpc/clients/cytology"

	"github.com/google/uuid"
)

var segTypeMap = map[domain.SegType]pb.SegType{
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

var groupTypeMap = map[domain.GroupType]pb.GroupType{
	domain.GroupTypeCE: pb.GroupType_GROUP_TYPE_CE,
	domain.GroupTypeCL: pb.GroupType_GROUP_TYPE_CL,
	domain.GroupTypeME: pb.GroupType_GROUP_TYPE_ME,
}

func (a *adapter) CreateSegmentationGroup(ctx context.Context, in CreateSegmentationGroupIn) (int, error) {
	req := &pb.CreateSegmentationGroupIn{
		CytologyId: in.CytologyID.String(),
		SegType:    segTypeMap[in.SegType],
		GroupType:  groupTypeMap[in.GroupType],
		IsAi:       in.IsAI,
		Details:    in.Details,
	}

	res, err := a.client.CreateSegmentationGroup(ctx, req)
	if err != nil {
		return 0, adapter_errors.HandleGRPCError(err)
	}

	return int(res.Id), nil
}

func (a *adapter) GetSegmentationGroupsByCytologyId(ctx context.Context, id uuid.UUID, segType *domain.SegType, groupType *domain.GroupType, isAI *bool) ([]domain.SegmentationGroup, error) {
	req := &pb.GetSegmentationGroupsByCytologyIdIn{
		CytologyId: id.String(),
	}

	if segType != nil {
		st := segTypeMap[*segType]
		req.SegType = &st
	}

	if groupType != nil {
		gt := groupTypeMap[*groupType]
		req.GroupType = &gt
	}

	if isAI != nil {
		req.IsAi = isAI
	}

	res, err := a.client.GetSegmentationGroupsByCytologyId(ctx, req)
	if err != nil {
		return nil, adapter_errors.HandleGRPCError(err)
	}

	return mappers.SegmentationGroup{}.SliceDomain(res.SegmentationGroups), nil
}

func (a *adapter) UpdateSegmentationGroup(ctx context.Context, in UpdateSegmentationGroupIn) (domain.SegmentationGroup, error) {
	req := &pb.UpdateSegmentationGroupIn{
		Id:      int32(in.Id),
		Details: in.Details,
	}

	if in.SegType != nil {
		st := segTypeMap[*in.SegType]
		req.SegType = &st
	}

	res, err := a.client.UpdateSegmentationGroup(ctx, req)
	if err != nil {
		return domain.SegmentationGroup{}, adapter_errors.HandleGRPCError(err)
	}

	return mappers.SegmentationGroup{}.Domain(res.SegmentationGroup), nil
}

func (a *adapter) DeleteSegmentationGroup(ctx context.Context, id int) error {
	_, err := a.client.DeleteSegmentationGroup(ctx, &pb.DeleteSegmentationGroupIn{Id: int32(id)})
	return adapter_errors.HandleGRPCError(err)
}

func (a *adapter) CreateSegmentation(ctx context.Context, in CreateSegmentationIn) (int, error) {
	points := make([]*pb.SegmentationPointCreate, 0, len(in.Points))
	for _, p := range in.Points {
		points = append(points, &pb.SegmentationPointCreate{
			X: int32(p.X),
			Y: int32(p.Y),
		})
	}

	req := &pb.CreateSegmentationIn{
		SegmentationGroupId: int32(in.SegmentationGroupID),
		Points:              points,
	}

	res, err := a.client.CreateSegmentation(ctx, req)
	if err != nil {
		return 0, adapter_errors.HandleGRPCError(err)
	}

	return int(res.Id), nil
}

func (a *adapter) GetSegmentationById(ctx context.Context, id int) (domain.Segmentation, error) {
	res, err := a.client.GetSegmentationById(ctx, &pb.GetSegmentationByIdIn{Id: int32(id)})
	if err != nil {
		return domain.Segmentation{}, adapter_errors.HandleGRPCError(err)
	}

	return mappers.Segmentation{}.Domain(res.Segmentation), nil
}

func (a *adapter) GetSegmentsByGroupId(ctx context.Context, id int) ([]domain.Segmentation, error) {
	res, err := a.client.GetSegmentsByGroupId(ctx, &pb.GetSegmentsByGroupIdIn{SegmentationGroupId: int32(id)})
	if err != nil {
		return nil, adapter_errors.HandleGRPCError(err)
	}

	return mappers.Segmentation{}.SliceDomain(res.Segmentations), nil
}

func (a *adapter) UpdateSegmentation(ctx context.Context, in UpdateSegmentationIn) (domain.Segmentation, error) {
	points := make([]*pb.SegmentationPointCreate, 0, len(in.Points))
	for _, p := range in.Points {
		points = append(points, &pb.SegmentationPointCreate{
			X: int32(p.X),
			Y: int32(p.Y),
		})
	}

	req := &pb.UpdateSegmentationIn{
		Id:     int32(in.Id),
		Points: points,
	}

	res, err := a.client.UpdateSegmentation(ctx, req)
	if err != nil {
		return domain.Segmentation{}, adapter_errors.HandleGRPCError(err)
	}

	return mappers.Segmentation{}.Domain(res.Segmentation), nil
}

func (a *adapter) DeleteSegmentation(ctx context.Context, id int) error {
	_, err := a.client.DeleteSegmentation(ctx, &pb.DeleteSegmentationIn{Id: int32(id)})
	return adapter_errors.HandleGRPCError(err)
}
