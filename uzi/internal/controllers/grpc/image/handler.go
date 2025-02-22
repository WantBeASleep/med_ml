package image

import (
	"context"

	pb "uzi/internal/generated/grpc/service"
	"uzi/internal/services/image"



	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type ImageHandler interface {
	GetUziImages(ctx context.Context, in *pb.GetUziImagesIn) (*pb.GetUziImagesOut, error)
}

type handler struct {
	imageSrv image.Service
}

func New(
	imageSrv image.Service,
) ImageHandler {
	return &handler{
		imageSrv: imageSrv,
	}
}

func (h *handler) GetUziImages(ctx context.Context, in *pb.GetUziImagesIn) (*pb.GetUziImagesOut, error) {
	images, err := h.imageSrv.GetUziImages(ctx, uuid.MustParse(in.UziId))
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Что то пошло не так: %s", err.Error())
	}

	out := pb.GetUziImagesOut{}
	for _, v := range images {
		out.Images = append(out.Images, &pb.Image{
			Id:   v.Id.String(),
			Page: int64(v.Page),
		})
	}

	return &out, nil
}


