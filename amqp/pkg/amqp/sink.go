package amqp

import (
	"context"
	"fmt"
	"github.com/cloudevents/sdk-go"
	"github.com/n3wscott/kpax/amqp/pkg/adapters"
	"github.com/n3wscott/kpax/amqp/pkg/bridge"
)

func NewSink(opt Options) adapters.Adapter {
	return &sink{Options: opt}
}

type sink struct {
	Options
}

func (a *sink) Start(ctx context.Context) error {
	inbound, err := cloudevents.NewDefaultClient()
	if err != nil {
		return fmt.Errorf("sink failed to create outbound client, %s", err.Error())
	}

	outbound, err := makeAMQPClient(a.Options)
	if err != nil {
		return fmt.Errorf("sink failed to create inbound client, %s", err.Error())
	}

	return bridge.Bridge(ctx, inbound, outbound)
}
