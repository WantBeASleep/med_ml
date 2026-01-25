package cytology

import (
	"context"

	"github.com/google/uuid"

	"composition-api/internal/adapters/cytology"
	domain "composition-api/internal/domain/cytology"
)

func (s *service) CreateSegmentationGroup(ctx context.Context, arg CreateSegmentationGroupArg) (int, error) {
	return s.adapters.Cytology.CreateSegmentationGroup(ctx, cytology.CreateSegmentationGroupIn{
		CytologyID: arg.CytologyID,
		SegType:    arg.SegType,
		GroupType:  arg.GroupType,
		IsAI:       arg.IsAI,
		Details:    arg.Details,
	})
}

func (s *service) GetSegmentationGroupsByCytologyId(ctx context.Context, id uuid.UUID, segType *domain.SegType, groupType *domain.GroupType, isAI *bool) ([]domain.SegmentationGroup, error) {
	return s.adapters.Cytology.GetSegmentationGroupsByCytologyId(ctx, id, segType, groupType, isAI)
}

func (s *service) UpdateSegmentationGroup(ctx context.Context, arg UpdateSegmentationGroupArg) (domain.SegmentationGroup, error) {
	return s.adapters.Cytology.UpdateSegmentationGroup(ctx, cytology.UpdateSegmentationGroupIn{
		Id:      arg.Id,
		SegType: arg.SegType,
		Details: arg.Details,
	})
}

func (s *service) DeleteSegmentationGroup(ctx context.Context, id int) error {
	return s.adapters.Cytology.DeleteSegmentationGroup(ctx, id)
}

func (s *service) CreateSegmentation(ctx context.Context, arg CreateSegmentationArg) (int, error) {
	return s.adapters.Cytology.CreateSegmentation(ctx, cytology.CreateSegmentationIn{
		SegmentationGroupID: arg.SegmentationGroupID,
		Points:              arg.Points,
	})
}

func (s *service) GetSegmentationById(ctx context.Context, id int) (domain.Segmentation, error) {
	return s.adapters.Cytology.GetSegmentationById(ctx, id)
}

func (s *service) GetSegmentsByGroupId(ctx context.Context, id int) ([]domain.Segmentation, error) {
	return s.adapters.Cytology.GetSegmentsByGroupId(ctx, id)
}

func (s *service) UpdateSegmentation(ctx context.Context, arg UpdateSegmentationArg) (domain.Segmentation, error) {
	return s.adapters.Cytology.UpdateSegmentation(ctx, cytology.UpdateSegmentationIn{
		Id:     arg.Id,
		Points: arg.Points,
	})
}

func (s *service) DeleteSegmentation(ctx context.Context, id int) error {
	return s.adapters.Cytology.DeleteSegmentation(ctx, id)
}
