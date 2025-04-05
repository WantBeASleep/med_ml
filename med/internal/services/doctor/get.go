package doctor

import (
	"context"
	"fmt"

	"med/internal/domain"

	"github.com/google/uuid"
)

func (s *service) GetDoctor(ctx context.Context, id uuid.UUID) (domain.Doctor, error) {
	doctor, err := s.dao.NewDoctorQuery(ctx).GetDoctorByID(id)
	if err != nil {
		return domain.Doctor{}, fmt.Errorf("get doctor by pk: %w", err)
	}

	return doctor.ToDomain(), nil
}
