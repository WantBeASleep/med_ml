package mappers

import (
	"med/internal/domain"
	pb "med/internal/generated/grpc/service"
)

func DoctorFromDomain(domain domain.Doctor) *pb.Doctor {
	return &pb.Doctor{
		Id:          domain.Id.String(),
		Fullname:    domain.FullName,
		Org:         domain.Org,
		Job:         domain.Job,
		Description: domain.Description,
	}
}
