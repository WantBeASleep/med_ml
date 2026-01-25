package mappers

import (
	"cytology/internal/domain"
	pb "cytology/internal/generated/grpc/service"
	"cytology/internal/services/cytology_image"

	"github.com/google/uuid"
)

var DiagnosticMarkingMap = map[domain.DiagnosticMarking]pb.DiagnosticMarking{
	domain.DiagnosticMarkingP11: pb.DiagnosticMarking_DIAGNOSTIC_MARKING_P11,
	domain.DiagnosticMarkingL23: pb.DiagnosticMarking_DIAGNOSTIC_MARKING_L23,
}

var DiagnosticMarkingReverseMap = map[pb.DiagnosticMarking]domain.DiagnosticMarking{
	pb.DiagnosticMarking_DIAGNOSTIC_MARKING_P11: domain.DiagnosticMarkingP11,
	pb.DiagnosticMarking_DIAGNOSTIC_MARKING_L23: domain.DiagnosticMarkingL23,
}

var MaterialTypeMap = map[domain.MaterialType]pb.MaterialType{
	domain.MaterialTypeGS:  pb.MaterialType_MATERIAL_TYPE_GS,
	domain.MaterialTypeBP:  pb.MaterialType_MATERIAL_TYPE_BP,
	domain.MaterialTypeTP:  pb.MaterialType_MATERIAL_TYPE_TP,
	domain.MaterialTypePTP: pb.MaterialType_MATERIAL_TYPE_PTP,
	domain.MaterialTypeLNP: pb.MaterialType_MATERIAL_TYPE_LNP,
}

var MaterialTypeReverseMap = map[pb.MaterialType]domain.MaterialType{
	pb.MaterialType_MATERIAL_TYPE_GS:  domain.MaterialTypeGS,
	pb.MaterialType_MATERIAL_TYPE_BP:  domain.MaterialTypeBP,
	pb.MaterialType_MATERIAL_TYPE_TP:  domain.MaterialTypeTP,
	pb.MaterialType_MATERIAL_TYPE_PTP: domain.MaterialTypePTP,
	pb.MaterialType_MATERIAL_TYPE_LNP: domain.MaterialTypeLNP,
}

func CreateCytologyImageArgFromProto(in *pb.CreateCytologyImageIn, externalID, doctorID, patientID uuid.UUID, prevID, parentPrevID *uuid.UUID) cytology_image.CreateCytologyImageArg {
	arg := cytology_image.CreateCytologyImageArg{
		ExternalID:       externalID,
		DoctorID:         doctorID,
		PatientID:        patientID,
		DiagnosticNumber: int(in.DiagnosticNumber),
		PrevID:           prevID,
		ParentPrevID:     parentPrevID,
	}

	if in.DiagnosticMarking != nil {
		dm := DiagnosticMarkingReverseMap[*in.DiagnosticMarking]
		arg.DiagnosticMarking = &dm
	}

	if in.MaterialType != nil {
		mt := MaterialTypeReverseMap[*in.MaterialType]
		arg.MaterialType = &mt
	}

	if in.Calcitonin != nil {
		calc := int(*in.Calcitonin)
		arg.Calcitonin = &calc
	}

	if in.CalcitoninInFlush != nil {
		calc := int(*in.CalcitoninInFlush)
		arg.CalcitoninInFlush = &calc
	}

	if in.Thyroglobulin != nil {
		thy := int(*in.Thyroglobulin)
		arg.Thyroglobulin = &thy
	}

	if in.Details != nil && *in.Details != "" {
		arg.Details = []byte(*in.Details)
	}

	if in.File != nil && len(in.File) > 0 {
		arg.File = in.File
	}

	if in.ContentType != nil && *in.ContentType != "" {
		arg.ContentType = *in.ContentType
	}

	return arg
}

func UpdateCytologyImageArgFromProto(in *pb.UpdateCytologyImageIn, id uuid.UUID) cytology_image.UpdateCytologyImageArg {
	arg := cytology_image.UpdateCytologyImageArg{
		Id: id,
	}

	if in.DiagnosticMarking != nil {
		dm := DiagnosticMarkingReverseMap[*in.DiagnosticMarking]
		arg.DiagnosticMarking = &dm
	}

	if in.MaterialType != nil {
		mt := MaterialTypeReverseMap[*in.MaterialType]
		arg.MaterialType = &mt
	}

	if in.Calcitonin != nil {
		calc := int(*in.Calcitonin)
		arg.Calcitonin = &calc
	}

	if in.CalcitoninInFlush != nil {
		calc := int(*in.CalcitoninInFlush)
		arg.CalcitoninInFlush = &calc
	}

	if in.Thyroglobulin != nil {
		thy := int(*in.Thyroglobulin)
		arg.Thyroglobulin = &thy
	}

	if in.Details != nil && *in.Details != "" {
		arg.Details = []byte(*in.Details)
	}

	if in.IsLast != nil {
		arg.IsLast = in.IsLast
	}

	if in.PrevId != nil && *in.PrevId != "" {
		parsed, err := uuid.Parse(*in.PrevId)
		if err == nil {
			arg.PrevID = &parsed
		}
	}

	if in.ParentPrevId != nil && *in.ParentPrevId != "" {
		parsed, err := uuid.Parse(*in.ParentPrevId)
		if err == nil {
			arg.ParentPrevID = &parsed
		}
	}

	return arg
}

func CytologyImageToProto(d domain.CytologyImage) *pb.CytologyImage {
	pbImg := &pb.CytologyImage{
		Id:               d.Id.String(),
		ExternalId:       d.ExternalID.String(),
		DoctorId:         d.DoctorID.String(),
		PatientId:        d.PatientID.String(),
		DiagnosticNumber: int32(d.DiagnosticNumber),
		IsLast:           d.IsLast,
		DiagnosDate:      d.DiagnosDate.Format("2006-01-02T15:04:05Z07:00"),
		CreateAt:         d.CreateAt.Format("2006-01-02T15:04:05Z07:00"),
	}

	if d.DiagnosticMarking != nil {
		dm := DiagnosticMarkingMap[*d.DiagnosticMarking]
		pbImg.DiagnosticMarking = &dm
	}

	if d.MaterialType != nil {
		mt := MaterialTypeMap[*d.MaterialType]
		pbImg.MaterialType = &mt
	}

	if d.Calcitonin != nil {
		calc := int32(*d.Calcitonin)
		pbImg.Calcitonin = &calc
	}

	if d.CalcitoninInFlush != nil {
		calc := int32(*d.CalcitoninInFlush)
		pbImg.CalcitoninInFlush = &calc
	}

	if d.Thyroglobulin != nil {
		thy := int32(*d.Thyroglobulin)
		pbImg.Thyroglobulin = &thy
	}

	if d.Details != nil {
		details := string(d.Details)
		pbImg.Details = &details
	}

	if d.PrevID != nil {
		prev := d.PrevID.String()
		pbImg.PrevId = &prev
	}

	if d.ParentPrevID != nil {
		parent := d.ParentPrevID.String()
		pbImg.ParentPrevId = &parent
	}

	return pbImg
}

func OriginalImageToProto(d domain.OriginalImage) *pb.OriginalImage {
	pbImg := &pb.OriginalImage{
		Id:         d.Id.String(),
		CytologyId: d.CytologyID.String(),
		ImagePath:  d.ImagePath,
		CreateDate: d.CreateDate.Format("2006-01-02T15:04:05Z07:00"),
		ViewedFlag: d.ViewedFlag,
	}

	if d.DelayTime != nil {
		delay := *d.DelayTime
		pbImg.DelayTime = &delay
	}

	return pbImg
}
