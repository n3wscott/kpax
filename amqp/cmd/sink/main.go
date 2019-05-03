package main

import (
	"context"
	"fmt"
	"github.com/n3wscott/kpax/amqp/pkg/amqp"
	"log"
	"os"

	"github.com/kelseyhightower/envconfig"
)

type envConfig struct {
	AMQPServer    string `envconfig:"AMQP_SERVER" default:"amqp://localhost:5672/" required:"true"`
	Queue         string `envconfig:"AMQP_QUEUE" required:"true"`
	AccessKeyName string `envconfig:"AMQP_ACCESS_KEY_NAME" default:"guest" required:"true"`
	AccessKey     string `envconfig:"AMQP_ACCESS_KEY" default:"password" required:"true"`
}

func main() {
	var env envConfig
	if err := envconfig.Process("", &env); err != nil {
		log.Printf("[ERROR] Failed to process env var: %s", err)
		os.Exit(1)
	}

	ra := amqp.NewSink(amqp.Options{
		AMQPServer:    env.AMQPServer,
		Queue:         env.Queue,
		AccessKey:     env.AccessKey,
		AccessKeyName: env.AccessKeyName,
	})

	if err := ra.Start(context.Background()); err != nil {
		fmt.Printf("sink returned an error, %s\n", err)
		os.Exit(1)
	}

	os.Exit(0)
}
