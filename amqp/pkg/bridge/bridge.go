package bridge

import (
	"context"
	cloudevents "github.com/cloudevents/sdk-go"
)

// This bridges a CloudEvents Inbound Client to a CloudEvents outbound client.

type adapter struct {
	client cloudevents.Client
}

func Bridge(ctx context.Context, inbound, outbound cloudevents.Client) error {
	a := adapter{client: outbound}
	return inbound.StartReceiver(ctx, a.Receive)
}

func (a *adapter) Receive(ctx context.Context, event cloudevents.Event, resp *cloudevents.EventResponse) error {
	ret, err := a.client.Send(ctx, event)
	if err != nil {
		return err
	}
	if ret != nil {
		resp.RespondWith(200, ret)
	}
	return nil
}
