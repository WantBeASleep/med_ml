package doctor

import (
	"context"
	"fmt"

	"med/internal/domain"
	dentity "med/internal/repository/doctor/entity"
)

func (s *service) RegisterDoctor(ctx context.Context, doctor domain.Doctor) error {
	if err := s.dao.NewDoctorQuery(ctx).InsertDoctor(dentity.Doctor{}.FromDomain(doctor)); err != nil {
		return fmt.Errorf("insert doctor: %w", err)
	}

	return nil
}
