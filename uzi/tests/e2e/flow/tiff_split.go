package flow

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/IBM/sarama"
	"github.com/google/uuid"
	minio "github.com/minio/minio-go/v7"
	"google.golang.org/protobuf/proto"

	"uzi/internal/domain"
	pbDbus "uzi/internal/generated/dbus/consume/uziupload"
	pbSrv "uzi/internal/generated/grpc/service"
)

var TiffSplit flowfuncDepsInjector = func(deps *Deps) flowfunc {
	return func(ctx context.Context, data FlowData) (FlowData, error) {
		tiffFile, err := os.Open("assets/sample.tiff")
		if err != nil {
			return FlowData{}, fmt.Errorf("open tiff file: %w", err)
		}
		defer tiffFile.Close()

		fileInfo, err := tiffFile.Stat()
		if err != nil {
			return FlowData{}, fmt.Errorf("get file info: %w", err)
		}

		_, err = deps.S3.PutObject(
			ctx,
			deps.Bucket,
			filepath.Join(data.Uzi.Id.String(), data.Uzi.Id.String()),
			tiffFile,
			fileInfo.Size(),
			minio.PutObjectOptions{
				ContentType: "image/tiff",
			},
		)
		if err != nil {
			return FlowData{}, fmt.Errorf("put object: %w", err)
		}

		// тригерим событие обработки загруженного узи
		message, err := proto.Marshal(&pbDbus.UziUpload{UziId: data.Uzi.Id.String()})
		if err != nil {
			return FlowData{}, fmt.Errorf("marshal message: %w", err)
		}

		_, _, err = deps.Dbus.SendMessage(
			&sarama.ProducerMessage{
				Topic: "uziupload",
				Value: sarama.ByteEncoder(message),
			},
		)
		if err != nil {
			return FlowData{}, fmt.Errorf("send message: %w", err)
		}

		// ретраимся пока не получим статус обработанного узи
		ctx, cancel := context.WithTimeout(ctx, time.Second*10)
		defer cancel()

		backoff := time.Second * 1
	uzi_status_waiter:
		for {
			select {
			case <-ctx.Done():
				return FlowData{}, fmt.Errorf("context done. uzi not splitted")

			case <-time.After(backoff):
				resp, err := deps.Adapter.GetUziById(ctx, &pbSrv.GetUziByIdIn{Id: data.Uzi.Id.String()})
				if err != nil {
					return FlowData{}, fmt.Errorf("get uzi by id: %w", err)
				}

				if resp.Uzi.Status != pbSrv.UziStatus_UZI_STATUS_PENDING {
					backoff *= 2
					continue
				}

				break uzi_status_waiter
			}
		}

		resp, err := deps.Adapter.GetImagesByUziId(ctx, &pbSrv.GetImagesByUziIdIn{UziId: data.Uzi.Id.String()})
		if err != nil {
			return FlowData{}, fmt.Errorf("get images by uzi id: %w", err)
		}

		for _, image := range resp.Images {
			data.Images = append(data.Images, domain.Image{
				Id:    uuid.MustParse(image.Id),
				UziID: uuid.MustParse(image.UziId),
				Page:  int(image.Page),
			})
		}

		return data, nil
	}
}
