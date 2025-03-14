package uzi

import (
	"context"
	"fmt"
	"os"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/google/uuid"
	ht "github.com/ogen-go/ogen/http"

	"gateway/internal/generated/http/api"
)

var UziInit flowfuncDepsInjector = func(deps *Deps) flowfunc {
	return func(ctx context.Context, data FlowData) (FlowData, error) {
		tiffFile, err := os.Open(os.Getenv("ASSETS_PATH"))
		if err != nil {
			return FlowData{}, fmt.Errorf("open tiff file: %w", err)
		}
		defer tiffFile.Close()

		fileInfo, err := tiffFile.Stat()
		if err != nil {
			return FlowData{}, fmt.Errorf("get file info: %w", err)
		}

		file := ht.MultipartFile{
			Name: fileInfo.Name(),
			File: tiffFile,
			Size: fileInfo.Size(),
		}

		projection := gofakeit.Word()
		externalId := uuid.New()

		resp, err := deps.Adapter.UziPost(ctx, &api.UziPostReq{
			File:       file,
			Projection: projection,
			ExternalID: externalId,
			DeviceID:   data.Got.DeviceID,
		})
		if err != nil {
			return FlowData{}, fmt.Errorf("create uzi: %w", err)
		}

		switch v := resp.(type) {
		case *api.SimpleUuid:
			data.Got.UziID = v.ID
			data.Expected.UziProjection = projection
			data.Expected.UziExternalID = externalId

		case *api.ErrorStatusCode:
			return FlowData{}, fmt.Errorf("create uzi error: %w", v)

		default:
			return FlowData{}, fmt.Errorf("unexpected uzi post response: %T", v)
		}

		return data, nil
	}
}
