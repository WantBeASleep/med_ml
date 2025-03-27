package uzi

import (
	"context"
	"fmt"
	"time"

	domain "composition-api/internal/domain/uzi"
	"composition-api/internal/generated/http/api"
)

var SaveImages flowfuncDepsInjector = func(deps *Deps) flowfunc {
	return func(ctx context.Context, data FlowData) (FlowData, error) {
		// ретраимся пока не получим статус обработанного узи
		ctx, cancel := context.WithTimeout(ctx, time.Second*20)
		defer cancel()

		backoff := time.Second * 1
	uzi_status_waiter:
		for {
			select {
			case <-ctx.Done():
				return FlowData{}, fmt.Errorf("context done. uzi not splitted")

			case <-time.After(backoff):
				resp, err := deps.Adapter.UziIDGet(ctx, api.UziIDGetParams{ID: data.Got.UziID})
				if err != nil {
					return FlowData{}, fmt.Errorf("get uzi by id: %w", err)
				}

				var status api.UziStatus
				switch v := resp.(type) {
				case *api.Uzi:
					status = v.Status

				case *api.ErrorStatusCode:
					return FlowData{}, fmt.Errorf("get uzi by id: %w", v)

				default:
					return FlowData{}, fmt.Errorf("unexpected uzi get response: %T", v)
				}

				// pending && completed ==> обработано
				if status == api.UziStatusNew {
					backoff *= 2
					continue
				}

				break uzi_status_waiter
			}
		}

		resp, err := deps.Adapter.UziIDImagesGet(ctx, api.UziIDImagesGetParams{ID: data.Got.UziID})
		if err != nil {
			return FlowData{}, fmt.Errorf("get images by uzi id: %w", err)
		}

		switch v := resp.(type) {
		case *api.UziIDImagesGetOKApplicationJSON:
			for _, image := range *v {
				data.Got.Images = append(data.Got.Images, domain.Image{
					Id:    image.ID,
					UziID: image.UziID,
					Page:  int(image.Page),
				})
			}

		case *api.ErrorStatusCode:
			return FlowData{}, fmt.Errorf("get images by uzi id: %w", v)

		default:
			return FlowData{}, fmt.Errorf("unexpected uzi images get response: %T", v)
		}

		return data, nil
	}
}
