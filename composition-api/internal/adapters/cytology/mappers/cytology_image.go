package mappers

import (
	"time"

	"github.com/google/uuid"

	domain "composition-api/internal/domain/cytology"
	pb "composition-api/internal/generated/grpc/clients/cytology"
)

var diagnosticMarkingMap = map[pb.DiagnosticMarking]domain.DiagnosticMarking{
	pb.DiagnosticMarking_DIAGNOSTIC_MARKING_P11: domain.DiagnosticMarkingP11,
	pb.DiagnosticMarking_DIAGNOSTIC_MARKING_L23: domain.DiagnosticMarkingL23,
}

var materialTypeMap = map[pb.MaterialType]domain.MaterialType{
	pb.MaterialType_MATERIAL_TYPE_GS:  domain.MaterialTypeGS,
	pb.MaterialType_MATERIAL_TYPE_BP:  domain.MaterialTypeBP,
	pb.MaterialType_MATERIAL_TYPE_TP:  domain.MaterialTypeTP,
	pb.MaterialType_MATERIAL_TYPE_PTP: domain.MaterialTypePTP,
	pb.MaterialType_MATERIAL_TYPE_LNP: domain.MaterialTypeLNP,
}

type CytologyImage struct{}

func (m CytologyImage) Domain(pb *pb.CytologyImage) domain.CytologyImage {
	diagnosDate, _ := time.Parse(time.RFC3339, pb.DiagnosDate)
	createAt, _ := time.Parse(time.RFC3339, pb.CreateAt)

	var diagnosticMarking *domain.DiagnosticMarking
	if pb.DiagnosticMarking != nil {
		dm := diagnosticMarkingMap[*pb.DiagnosticMarking]
		diagnosticMarking = &dm
	}

	var materialType *domain.MaterialType
	if pb.MaterialType != nil {
		mt := materialTypeMap[*pb.MaterialType]
		materialType = &mt
	}

	var calcitonin *int
	if pb.Calcitonin != nil {
		c := int(*pb.Calcitonin)
		calcitonin = &c
	}

	var calcitoninInFlush *int
	if pb.CalcitoninInFlush != nil {
		c := int(*pb.CalcitoninInFlush)
		calcitoninInFlush = &c
	}

	var thyroglobulin *int
	if pb.Thyroglobulin != nil {
		t := int(*pb.Thyroglobulin)
		thyroglobulin = &t
	}

	var prevID *uuid.UUID
	if pb.PrevId != nil && *pb.PrevId != "" {
		prev := uuid.MustParse(*pb.PrevId)
		prevID = &prev
	}

	var parentPrevID *uuid.UUID
	if pb.ParentPrevId != nil && *pb.ParentPrevId != "" {
		parent := uuid.MustParse(*pb.ParentPrevId)
		parentPrevID = &parent
	}

	return domain.CytologyImage{
		Id:                uuid.MustParse(pb.Id),
		ExternalID:        uuid.MustParse(pb.ExternalId),
		DoctorID:          uuid.MustParse(pb.GetDoctorId()),
		PatientID:         uuid.MustParse(pb.GetPatientId()),
		DiagnosticNumber:  int(pb.DiagnosticNumber),
		DiagnosticMarking: diagnosticMarking,
		MaterialType:      materialType,
		DiagnosDate:       diagnosDate,
		IsLast:            pb.IsLast,
		Calcitonin:        calcitonin,
		CalcitoninInFlush: calcitoninInFlush,
		Thyroglobulin:     thyroglobulin,
		Details:           pb.Details,
		PrevID:            prevID,
		ParentPrevID:      parentPrevID,
		CreateAt:          createAt,
	}
}

func (m CytologyImage) SliceDomain(pbs []*pb.CytologyImage) []domain.CytologyImage {
	domains := make([]domain.CytologyImage, 0, len(pbs))
	for _, pb := range pbs {
		domains = append(domains, m.Domain(pb))
	}
	return domains
}
