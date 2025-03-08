package flow

import (
	"context"
	"fmt"
	"math/rand"
	"time"

	"github.com/IBM/sarama"
	"github.com/google/uuid"
	"google.golang.org/protobuf/proto"

	"uzi/internal/domain"
	pbDbus "uzi/internal/generated/dbus/consume/uziprocessed"
	pbSrv "uzi/internal/generated/grpc/service"
)

var SaveNodesWithSegments flowfuncDepsInjector = func(deps *Deps) flowfunc {
	return func(ctx context.Context, data FlowData) (FlowData, error) {
		request := &pbDbus.UziProcessed{
			UziId: data.Uzi.Id.String(),
		}

		// создадим 10 узлов
		for range 10 {
			nodeWithSegments := &pbDbus.NodeWithSegments{}

			node := &pbDbus.Node{
				Ai:        true,
				Tirads_23: rand.Float64(),
				Tirads_4:  rand.Float64(),
				Tirads_5:  rand.Float64(),
			}
			nodeWithSegments.Node = node

			// по 3 сегмента на узел
			for range 3 {
				imageId := data.Images[rand.Intn(len(data.Images))].Id
				segment := &pbDbus.Segment{
					ImageId:   imageId.String(),
					Contor:    []byte(`{[{"x": 1, "y": 1}]}`),
					Tirads_23: rand.Float64(),
					Tirads_4:  rand.Float64(),
					Tirads_5:  rand.Float64(),
				}

				nodeWithSegments.Segments = append(nodeWithSegments.Segments, segment)
			}

			request.NodesWithSegments = append(request.NodesWithSegments, nodeWithSegments)
		}

		message, err := proto.Marshal(request)
		if err != nil {
			return FlowData{}, fmt.Errorf("marshal request: %w", err)
		}

		_, _, err = deps.Dbus.SendMessage(
			&sarama.ProducerMessage{
				Topic: "uziprocessed",
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

				if resp.Uzi.Status != pbSrv.UziStatus_UZI_STATUS_COMPLETED {
					backoff *= 2
					continue
				}

				break uzi_status_waiter
			}
		}

		resp, err := deps.Adapter.GetNodesByUziId(ctx, &pbSrv.GetNodesByUziIdIn{UziId: data.Uzi.Id.String()})
		if err != nil {
			return FlowData{}, fmt.Errorf("get nodes by uzi id: %w", err)
		}

		for _, node := range resp.Nodes {
			data.Nodes = append(data.Nodes, domain.Node{
				Id:       uuid.MustParse(node.Id),
				Ai:       node.Ai,
				UziID:    data.Uzi.Id,
				Tirads23: node.Tirads_23,
				Tirads4:  node.Tirads_4,
				Tirads5:  node.Tirads_5,
			})
		}

		for _, node := range resp.Nodes {
			resp, err := deps.Adapter.GetSegmentsByNodeId(ctx, &pbSrv.GetSegmentsByNodeIdIn{NodeId: node.Id})
			if err != nil {
				return FlowData{}, fmt.Errorf("get segments by node id: %w", err)
			}

			for _, segment := range resp.Segments {
				data.Segments = append(data.Segments, domain.Segment{
					Id:       uuid.MustParse(segment.Id),
					ImageID:  uuid.MustParse(segment.ImageId),
					NodeID:   uuid.MustParse(segment.NodeId),
					Contor:   segment.Contor,
					Tirads23: segment.Tirads_23,
					Tirads4:  segment.Tirads_4,
					Tirads5:  segment.Tirads_5,
				})
			}
		}

		return data, nil
	}
}
