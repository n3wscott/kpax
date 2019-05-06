package amqp

import (
	"context"
	"fmt"
	"github.com/cloudevents/sdk-go"
	"github.com/n3wscott/kpax/amqp/pkg/adapters"
	"github.com/n3wscott/kpax/amqp/pkg/bridge"
	"strings"
)

func NewSink(opt Options) adapters.Adapter {
	return &sink{Options: opt}
}

type sink struct {
	Options
}

func (a *sink) Start(ctx context.Context) error {

	// Sink will listen on inbound HTTP.
	inbound, err := cloudevents.NewDefaultClient()
	if err != nil {
		return fmt.Errorf("sink failed to create inbound client, %s", err.Error())
	}

	// Sink will send outbound on AMQP.
	outbound, err := makeAMQPClient(a.Options)
	if err != nil {
		return fmt.Errorf("sink failed to create outbound client, %s", err.Error())
	}

	return bridge.Bridge(ctx, inbound, outbound, func(event *cloudevents.Event) bool {
		for k, v := range event.Extensions() {
			if strings.EqualFold(a.SinkAccessKeyName, k) {
				if strings.EqualFold(a.SinkAccessKey, v.(string)) {
					event.SetExtension(a.SinkAccessKeyName, nil) // delete it.
					return true
				}
			}
		}
		return false
	})
}
