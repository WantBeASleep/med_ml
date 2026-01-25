package original_image

import (
	"context"

	pb "cytology/internal/generated/grpc/service"
	"cytology/internal/services/original_image"

	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (h *handler) CreateOriginalImage(ctx context.Context, in *pb.CreateOriginalImageIn) (*pb.CreateOriginalImageOut, error) {
	cytologyID, err := uuid.Parse(in.CytologyId)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "cytology_id is not a valid uuid: %s", err.Error())
	}

	// Если передан путь к файлу, используем его (файл уже загружен в S3)
	// Иначе проверяем, что передан файл
	imagePath := in.GetImagePath()
	var imagePathPtr *string
	if imagePath != "" {
		imagePathPtr = &imagePath
	}

	if imagePathPtr == nil || *imagePathPtr == "" {
		if len(in.File) == 0 {
			return nil, status.Errorf(codes.InvalidArgument, "file or image_path is required")
		}
		if in.ContentType == "" {
			return nil, status.Errorf(codes.InvalidArgument, "content_type is required")
		}
	}

	var delayTimePtr *float64
	if in.DelayTime != nil {
		delayTime := in.GetDelayTime()
		delayTimePtr = &delayTime
	}

	arg := original_image.CreateOriginalImageArg{
		CytologyID:  cytologyID,
		File:        in.File,
		ContentType: in.ContentType,
		DelayTime:   delayTimePtr,
		ImagePath:   imagePathPtr,
	}

	id, err := h.services.OriginalImage.CreateOriginalImage(ctx, arg)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Что то пошло не так: %s", err.Error())
	}

	return &pb.CreateOriginalImageOut{Id: id.String()}, nil
}
