package amqp

import (
	"context"
	"fmt"
	"github.com/n3wscott/kpax/amqp/pkg/adapters"
	"github.com/n3wscott/kpax/amqp/pkg/bridge"
)

func NewSource(opt Options) adapters.Adapter {
	return &source{Options: opt}
}

type source struct {
	Options
}

func (a *source) Start(ctx context.Context) error {
	inbound, err := makeAMQPClient(a.Options)
	if err != nil {
		return fmt.Errorf("source failed to create inbound client, %s", err.Error())
	}

	outbound, err := makeHTTPClient(a.Options)
	if err != nil {
		return fmt.Errorf("source failed to create outbound client, %s", err.Error())
	}

	return bridge.Bridge(ctx, inbound, outbound, nil)
}
