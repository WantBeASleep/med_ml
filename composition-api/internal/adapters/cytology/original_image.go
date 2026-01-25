package cytology

import (
	"context"
	"io"

	"composition-api/internal/adapters/cytology/mappers"
	adapter_errors "composition-api/internal/adapters/errors"
	domain "composition-api/internal/domain/cytology"
	pb "composition-api/internal/generated/grpc/clients/cytology"

	"github.com/google/uuid"
)

func (a *adapter) CreateOriginalImage(ctx context.Context, in CreateOriginalImageIn) (uuid.UUID, error) {
	req := &pb.CreateOriginalImageIn{
		CytologyId:  in.CytologyID.String(),
		ContentType: in.ContentType,
		DelayTime:   in.DelayTime,
	}

	// Если передан путь к файлу, используем его (файл уже загружен в S3)
	if in.ImagePath != nil && *in.ImagePath != "" {
		// Файл уже загружен в S3 в composition-api
		// Передаем путь вместо файла, чтобы не отправлять файл по сети
		req.ImagePath = in.ImagePath
		req.File = nil // Не передаем файл, так как он уже в S3
	} else {
		// Если путь не передан, читаем файл и передаем его (старый способ)
		fileData, err := io.ReadAll(in.File.File)
		if err != nil {
			return uuid.Nil, adapter_errors.HandleGRPCError(err)
		}
		req.File = fileData
	}

	res, err := a.client.CreateOriginalImage(ctx, req)
	if err != nil {
		return uuid.Nil, adapter_errors.HandleGRPCError(err)
	}

	return uuid.MustParse(res.Id), nil
}

func (a *adapter) GetOriginalImageById(ctx context.Context, id uuid.UUID) (domain.OriginalImage, error) {
	res, err := a.client.GetOriginalImageById(ctx, &pb.GetOriginalImageByIdIn{Id: id.String()})
	if err != nil {
		return domain.OriginalImage{}, adapter_errors.HandleGRPCError(err)
	}

	return mappers.OriginalImage{}.Domain(res.OriginalImage), nil
}

func (a *adapter) GetOriginalImagesByCytologyId(ctx context.Context, id uuid.UUID) ([]domain.OriginalImage, error) {
	res, err := a.client.GetOriginalImagesByCytologyId(ctx, &pb.GetOriginalImagesByCytologyIdIn{CytologyId: id.String()})
	if err != nil {
		return nil, adapter_errors.HandleGRPCError(err)
	}

	return mappers.OriginalImage{}.SliceDomain(res.OriginalImages), nil
}

func (a *adapter) UpdateOriginalImage(ctx context.Context, in UpdateOriginalImageIn) (domain.OriginalImage, error) {
	req := &pb.UpdateOriginalImageIn{
		Id:         in.Id.String(),
		DelayTime:  in.DelayTime,
		ViewedFlag: in.ViewedFlag,
	}

	res, err := a.client.UpdateOriginalImage(ctx, req)
	if err != nil {
		return domain.OriginalImage{}, adapter_errors.HandleGRPCError(err)
	}

	return mappers.OriginalImage{}.Domain(res.OriginalImage), nil
}
