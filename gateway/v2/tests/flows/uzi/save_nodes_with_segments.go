package uzi

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	domain "gateway/internal/domain/uzi"
	"gateway/internal/generated/http/api"
)

var SaveNodesWithSegments flowfuncDepsInjector = func(deps *Deps) flowfunc {
	return func(ctx context.Context, data FlowData) (FlowData, error) {
		ctx, cancel := context.WithTimeout(ctx, time.Minute)
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

				if status != api.UziStatusCompleted {
					backoff *= 2
					continue
				}

				break uzi_status_waiter
			}
		}

		nodesResp, err := deps.Adapter.UziIDNodesGet(ctx, api.UziIDNodesGetParams{ID: data.Got.UziID})
		if err != nil {
			return FlowData{}, fmt.Errorf("get nodes by uzi id: %w", err)
		}

		switch v := nodesResp.(type) {
		case *api.UziIDNodesGetOKApplicationJSON:
			for _, node := range *v {
				data.Got.Nodes = append(data.Got.Nodes, domain.Node{
					Id:       node.ID,
					Ai:       node.Ai,
					UziID:    data.Got.UziID,
					Tirads23: node.Tirads23,
					Tirads4:  node.Tirads4,
					Tirads5:  node.Tirads5,
				})
			}

		case *api.ErrorStatusCode:
			return FlowData{}, fmt.Errorf("get nodes by uzi id: %w", v)

		default:
			return FlowData{}, fmt.Errorf("unexpected uzi nodes get response: %T", v)
		}

		for _, node := range data.Got.Nodes {
			resp, err := deps.Adapter.UziNodesIDSegmentsGet(ctx, api.UziNodesIDSegmentsGetParams{ID: node.Id})
			if err != nil {
				return FlowData{}, fmt.Errorf("get segments by node id: %w", err)
			}

			switch v := resp.(type) {
			case *api.UziNodesIDSegmentsGetOKApplicationJSON:
				for _, segment := range *v {
					contor, err := json.Marshal(segment.Contor)
					if err != nil {
						return FlowData{}, fmt.Errorf("marshal contor: %w", err)
					}

					data.Got.Segments = append(data.Got.Segments, domain.Segment{
						Id:       segment.ID,
						ImageID:  segment.ImageID,
						NodeID:   segment.NodeID,
						Contor:   contor,
						Tirads23: segment.Tirads23,
						Tirads4:  segment.Tirads4,
						Tirads5:  segment.Tirads5,
					})
				}

			case *api.ErrorStatusCode:
				return FlowData{}, fmt.Errorf("get segments by node id: %w", v)

			default:
				return FlowData{}, fmt.Errorf("unexpected uzi nodes segments get response: %T", v)
			}
		}

		return data, nil
	}
}
