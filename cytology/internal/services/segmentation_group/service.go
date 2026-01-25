package segmentation_group

import (
	"context"
	"encoding/json"
	"time"

	"cytology/internal/domain"
	"cytology/internal/repository"
	segmentation_groupEntity "cytology/internal/repository/segmentation_group/entity"

	"github.com/google/uuid"
)

type Service interface {
	CreateSegmentationGroup(ctx context.Context, arg CreateSegmentationGroupArg) (int, error)
	GetSegmentationGroupByID(ctx context.Context, id int) (domain.SegmentationGroup, error)
	GetSegmentationGroupsByCytologyID(ctx context.Context, cytologyID uuid.UUID) ([]domain.SegmentationGroup, error)
	UpdateSegmentationGroup(ctx context.Context, arg UpdateSegmentationGroupArg) (domain.SegmentationGroup, error)
	DeleteSegmentationGroup(ctx context.Context, id int) error
}

type service struct {
	dao repository.DAO
}

func New(dao repository.DAO) Service {
	return &service{
		dao: dao,
	}
}

type CreateSegmentationGroupArg struct {
	CytologyID uuid.UUID
	SegType    domain.SegType
	GroupType  domain.GroupType
	IsAI       bool
	Details    []byte
}

type UpdateSegmentationGroupArg struct {
	Id      int
	SegType *domain.SegType
	Details []byte
}

func (s *service) CreateSegmentationGroup(ctx context.Context, arg CreateSegmentationGroupArg) (int, error) {
	var details json.RawMessage
	if arg.Details != nil && len(arg.Details) > 0 {
		// Проверяем, что это валидный JSON
		if string(arg.Details) != "null" {
			details = arg.Details
		}
	}

	group := domain.SegmentationGroup{
		Id:         0, // ID будет сгенерирован БД
		CytologyID: arg.CytologyID,
		SegType:    arg.SegType,
		GroupType:  arg.GroupType,
		IsAI:       arg.IsAI,
		Details:    details,
		CreateAt:   time.Now(),
	}

	entityGroup := segmentation_groupEntity.SegmentationGroup{}.FromDomain(group)
	id, err := s.dao.NewSegmentationGroupQuery(ctx).InsertSegmentationGroup(entityGroup)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (s *service) GetSegmentationGroupByID(ctx context.Context, id int) (domain.SegmentationGroup, error) {
	group, err := s.dao.NewSegmentationGroupQuery(ctx).GetSegmentationGroupByID(id)
	if err != nil {
		return domain.SegmentationGroup{}, err
	}
	return group.ToDomain(), nil
}

func (s *service) GetSegmentationGroupsByCytologyID(ctx context.Context, cytologyID uuid.UUID) ([]domain.SegmentationGroup, error) {
	groups, err := s.dao.NewSegmentationGroupQuery(ctx).GetSegmentationGroupsByCytologyID(cytologyID)
	if err != nil {
		return nil, err
	}
	return segmentation_groupEntity.SegmentationGroup{}.SliceToDomain(groups), nil
}

func (s *service) UpdateSegmentationGroup(ctx context.Context, arg UpdateSegmentationGroupArg) (domain.SegmentationGroup, error) {
	group, err := s.dao.NewSegmentationGroupQuery(ctx).GetSegmentationGroupByID(arg.Id)
	if err != nil {
		return domain.SegmentationGroup{}, err
	}

	domainGroup := group.ToDomain()
	if arg.SegType != nil {
		domainGroup.SegType = *arg.SegType
	}
	if arg.Details != nil {
		domainGroup.Details = arg.Details
	}

	entityGroup := segmentation_groupEntity.SegmentationGroup{}.FromDomain(domainGroup)
	if err := s.dao.NewSegmentationGroupQuery(ctx).UpdateSegmentationGroup(entityGroup); err != nil {
		return domain.SegmentationGroup{}, err
	}

	return domainGroup, nil
}

func (s *service) DeleteSegmentationGroup(ctx context.Context, id int) error {
	return s.dao.NewSegmentationGroupQuery(ctx).DeleteSegmentationGroup(id)
}
