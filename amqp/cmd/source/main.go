package main

import (
	"context"
	"fmt"
	"github.com/n3wscott/kpax/amqp/pkg/amqp"
	"log"
	"net/url"
	"os"

	"github.com/kelseyhightower/envconfig"
)

type envConfig struct {
	AMQPServer    string `envconfig:"AMQP_SERVER" default:"amqp://localhost:5672/" required:"true"`
	Queue         string `envconfig:"AMQP_QUEUE" required:"true"`
	AccessKeyName string `envconfig:"AMQP_ACCESS_KEY_NAME" default:"guest" required:"true"`
	AccessKey     string `envconfig:"AMQP_ACCESS_KEY" default:"password" required:"true"`
	Sink          string `envconfig:"SINK" required:"true"`
}

func main() {
	var env envConfig
	if err := envconfig.Process("", &env); err != nil {
		log.Printf("[ERROR] Failed to process env var: %s", err)
		os.Exit(1)
	}

	sink, err := url.Parse(env.Sink)
	if err != nil {
		fmt.Printf("failed to parse sink url, %s\n", err)
		os.Exit(1)
	}

	ra := amqp.NewSource(amqp.Options{
		AMQPServer:    env.AMQPServer,
		Queue:         env.Queue,
		AccessKey:     env.AccessKey,
		AccessKeyName: env.AccessKeyName,
		Sink:          sink,
	})

	if err := ra.Start(context.Background()); err != nil {
		fmt.Printf("source returned an error, %s\n", err)
		os.Exit(1)
	}

	os.Exit(0)
}
