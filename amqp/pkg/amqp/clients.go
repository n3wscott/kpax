package amqp

import (
	"net/url"

	cloudevents "github.com/cloudevents/sdk-go"
	"github.com/cloudevents/sdk-go/pkg/cloudevents/transport/amqp"
	qp "pack.ag/amqp"
)

type Options struct {
	// AMQP:

	AMQPServer    string
	Queue         string
	AccessKeyName string
	AccessKey     string

	// HTTP:

	Sink *url.URL
}

func makeAMQPClient(opt Options) (cloudevents.Client, error) {
	t, err := amqp.New(opt.AMQPServer, opt.Queue,
		amqp.WithConnOpt(qp.ConnSASLPlain(opt.AccessKeyName, opt.AccessKey)),
	)
	if err != nil {
		return nil, err
	}

	c, err := cloudevents.NewClient(t, cloudevents.WithTimeNow(), cloudevents.WithUUIDs())
	if err != nil {
		return nil, err
	}

	return c, nil
}

func makeHTTPClient(opt Options) (cloudevents.Client, error) {
	t, err := cloudevents.NewHTTPTransport(
		cloudevents.WithTarget(opt.Sink.String()),
		cloudevents.WithBinaryEncoding(),
	)
	if err != nil {
		return nil, err
	}

	c, err := cloudevents.NewClient(t, cloudevents.WithTimeNow(), cloudevents.WithUUIDs())
	if err != nil {
		return nil, err
	}

	return c, nil
}
