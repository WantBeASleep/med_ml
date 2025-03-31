package mappers

import (
	"auth/internal/domain"
	pb "auth/internal/generated/grpc/service"
)

var RoleMap = map[domain.Role]pb.Role{
	domain.RoleDoctor:  pb.Role_ROLE_DOCTOR,
	domain.RolePatient: pb.Role_ROLE_PATIENT,
}

var RoleReversedMap = map[pb.Role]domain.Role{
	pb.Role_ROLE_DOCTOR:  domain.RoleDoctor,
	pb.Role_ROLE_PATIENT: domain.RolePatient,
}
