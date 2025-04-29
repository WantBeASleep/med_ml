package tariff_plan

import (
	"context"
	"fmt"

	"billing/internal/domain"
	"billing/internal/repository"
	"billing/internal/repository/entity"

	"github.com/google/uuid"
)

type Service interface {
	CreateTariffPlan(ctx context.Context, tariffPlan domain.TariffPlan) (uuid.UUID, error)
	GetTariffPlanByID(ctx context.Context, id uuid.UUID) (domain.TariffPlan, error)
	UpdateTariffPlan(ctx context.Context, tariffPlan domain.TariffPlan) error
	DeleteTariffPlan(ctx context.Context, id uuid.UUID) error
	ListTariffPlans(ctx context.Context) ([]domain.TariffPlan, error)
}

type service struct {
	dao repository.DAO
}

func New(dao repository.DAO) Service {
	return &service{
		dao: dao,
	}
}

func (s *service) CreateTariffPlan(ctx context.Context, tariffPlan domain.TariffPlan) (uuid.UUID, error) {
	tariffPlan.ID = uuid.New()
	err := s.dao.NewTariffPlanQuery(ctx).InsertTariffPlan(entity.TariffPlan{}.FromDomain(tariffPlan))
	if err != nil {
		return uuid.Nil, fmt.Errorf("failed to create tariff plan: %w", err)
	}
	return tariffPlan.ID, nil
}

func (s *service) GetTariffPlanByID(ctx context.Context, id uuid.UUID) (domain.TariffPlan, error) {
	tariffPlanDB, err := s.dao.NewTariffPlanQuery(ctx).GetTariffPlanByID(id)
	if err != nil {
		return domain.TariffPlan{}, fmt.Errorf("failed to get tariff plan by id: %w", err)
	}
	return tariffPlanDB.ToDomain(), nil
}

func (s *service) UpdateTariffPlan(ctx context.Context, tariffPlan domain.TariffPlan) error {
	_, err := s.dao.NewTariffPlanQuery(ctx).UpdateTariffPlan(entity.TariffPlan{}.FromDomain(tariffPlan))
	if err != nil {
		return fmt.Errorf("failed to update tariff plan: %w", err)
	}
	return nil
}

func (s *service) DeleteTariffPlan(ctx context.Context, id uuid.UUID) error {
	err := s.dao.NewTariffPlanQuery(ctx).DeleteTariffPlan(id)
	if err != nil {
		return fmt.Errorf("failed to delete tariff plan: %w", err)
	}
	return nil
}

func (s *service) ListTariffPlans(ctx context.Context) ([]domain.TariffPlan, error) {
	tariffPlansDB, err := s.dao.NewTariffPlanQuery(ctx).ListTariffPlans()
	if err != nil {
		return nil, fmt.Errorf("failed to list tariff plans: %w", err)
	}

	var tariffPlans []domain.TariffPlan
	for _, tp := range tariffPlansDB {
		tariffPlans = append(tariffPlans, tp.ToDomain())
	}

	return tariffPlans, nil
}
