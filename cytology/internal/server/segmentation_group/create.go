package segmentation_group

import (
	"context"
	"strings"

	"cytology/internal/domain"
	pb "cytology/internal/generated/grpc/service"
	"cytology/internal/services/segmentation_group"

	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// segTypeToString преобразует protobuf SegType в короткое строковое значение
func segTypeToString(segType pb.SegType) string {
	// Убираем префикс "SEG_TYPE_" из имени enum
	str := segType.String()
	if strings.HasPrefix(str, "SEG_TYPE_") {
		return strings.TrimPrefix(str, "SEG_TYPE_")
	}
	// Если формат неожиданный, возвращаем как есть (но это не должно произойти)
	return str
}

// groupTypeToString преобразует protobuf GroupType в короткое строковое значение
func groupTypeToString(groupType pb.GroupType) string {
	// Убираем префикс "GROUP_TYPE_" из имени enum
	str := groupType.String()
	if strings.HasPrefix(str, "GROUP_TYPE_") {
		return strings.TrimPrefix(str, "GROUP_TYPE_")
	}
	// Если формат неожиданный, возвращаем как есть (но это не должно произойти)
	return str
}

func (h *handler) CreateSegmentationGroup(ctx context.Context, in *pb.CreateSegmentationGroupIn) (*pb.CreateSegmentationGroupOut, error) {
	cytologyID, err := uuid.Parse(in.CytologyId)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "cytology_id is not a valid uuid: %s", err.Error())
	}

	var details []byte
	if in.Details != nil && *in.Details != "" {
		details = []byte(*in.Details)
	}

	// Преобразуем protobuf enum в короткие строковые значения
	segTypeStr := segTypeToString(in.SegType)
	groupTypeStr := groupTypeToString(in.GroupType)

	arg := segmentation_group.CreateSegmentationGroupArg{
		CytologyID: cytologyID,
		SegType:    domain.SegType(segTypeStr),
		GroupType:  domain.GroupType(groupTypeStr),
		IsAI:       in.IsAi,
		Details:    details,
	}

	id, err := h.services.SegmentationGroup.CreateSegmentationGroup(ctx, arg)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Что то пошло не так: %s", err.Error())
	}

	return &pb.CreateSegmentationGroupOut{Id: int32(id)}, nil
}
