package segmentation

import (
	"context"
	"time"

	"cytology/internal/domain"
	"cytology/internal/repository"
	segmentationEntity "cytology/internal/repository/segmentation/entity"
)

type Service interface {
	CreateSegmentation(ctx context.Context, arg CreateSegmentationArg) (int, error)
	GetSegmentationByID(ctx context.Context, id int) (domain.Segmentation, error)
	GetSegmentsByGroupID(ctx context.Context, groupID int) ([]domain.Segmentation, error)
	UpdateSegmentation(ctx context.Context, arg UpdateSegmentationArg) (domain.Segmentation, error)
	DeleteSegmentation(ctx context.Context, id int) error
}

type service struct {
	dao repository.DAO
}

func New(dao repository.DAO) Service {
	return &service{
		dao: dao,
	}
}

type CreateSegmentationArg struct {
	SegmentationGroupID int
	Points              []domain.SegmentationPoint
}

type UpdateSegmentationArg struct {
	Id     int
	Points []domain.SegmentationPoint
}

func (s *service) CreateSegmentation(ctx context.Context, arg CreateSegmentationArg) (int, error) {
	seg := domain.Segmentation{
		Id:                  0, // ID будет сгенерирован БД
		SegmentationGroupID: arg.SegmentationGroupID,
		Points:              arg.Points,
		CreateAt:            time.Now(),
	}

	// ID для точек будут сгенерированы БД
	for i := range seg.Points {
		seg.Points[i].Id = 0             // ID будет сгенерирован БД
		seg.Points[i].SegmentationID = 0 // Будет установлен после создания сегментации
		seg.Points[i].CreateAt = time.Now()
	}

	entitySeg := segmentationEntity.Segmentation{}.FromDomain(seg)
	id, err := s.dao.NewSegmentationQuery(ctx).InsertSegmentation(entitySeg)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (s *service) GetSegmentationByID(ctx context.Context, id int) (domain.Segmentation, error) {
	seg, err := s.dao.NewSegmentationQuery(ctx).GetSegmentationByID(id)
	if err != nil {
		return domain.Segmentation{}, err
	}
	return seg.ToDomain(), nil
}

func (s *service) GetSegmentsByGroupID(ctx context.Context, groupID int) ([]domain.Segmentation, error) {
	segs, err := s.dao.NewSegmentationQuery(ctx).GetSegmentsByGroupID(groupID)
	if err != nil {
		return nil, err
	}
	return segmentationEntity.Segmentation{}.SliceToDomain(segs), nil
}

func (s *service) UpdateSegmentation(ctx context.Context, arg UpdateSegmentationArg) (domain.Segmentation, error) {
	seg, err := s.dao.NewSegmentationQuery(ctx).GetSegmentationByID(arg.Id)
	if err != nil {
		return domain.Segmentation{}, err
	}

	domainSeg := seg.ToDomain()
	domainSeg.Points = arg.Points

	// ID для точек будут сгенерированы БД при обновлении
	for i := range domainSeg.Points {
		domainSeg.Points[i].Id = 0 // ID будет сгенерирован БД
		domainSeg.Points[i].SegmentationID = domainSeg.Id
		domainSeg.Points[i].CreateAt = time.Now()
	}

	entitySeg := segmentationEntity.Segmentation{}.FromDomain(domainSeg)
	if err := s.dao.NewSegmentationQuery(ctx).UpdateSegmentation(entitySeg); err != nil {
		return domain.Segmentation{}, err
	}

	return domainSeg, nil
}

func (s *service) DeleteSegmentation(ctx context.Context, id int) error {
	return s.dao.NewSegmentationQuery(ctx).DeleteSegmentation(id)
}
