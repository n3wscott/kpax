package main

import (
	"context"
	cloudevents "github.com/cloudevents/sdk-go"
	"github.com/kelseyhightower/envconfig"
	"github.com/n3wscott/kpax/sockeye/pkg/controller"
	"log"
	"net/http"
	"os"
	"strings"
)

type envConfig struct {
	FilePath string `envconfig:"FILE_PATH" default:"/var/run/ko/" required:"true"`
}

func main() {
	var env envConfig
	if err := envconfig.Process("", &env); err != nil {
		log.Printf("[ERROR] Failed to process env var: %s", err)
		os.Exit(1)
	}
	if !strings.HasSuffix(env.FilePath, "/") {
		env.FilePath = env.FilePath + "/"
	}

	c := controller.New(env.FilePath)

	c.Mux().Handle("/static/", http.StripPrefix("/static/",
		http.FileServer(http.Dir(env.FilePath+"static"))))

	t, err := cloudevents.NewHTTPTransport(
		cloudevents.WithBinaryEncoding(),
		cloudevents.WithPath("/ce"), // hack hack
	)
	if err != nil {
		log.Fatalf("failed to create cloudevents transport, %s", err.Error())
	}
	// I am doing this to allow root to be both POST for cloudevents and GET as root ui.
	c.Mux().HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			t.ServeHTTP(w, r)
			return
		}
		c.RootHandler(w, r)
	})
	t.Handler = c.Mux()

	ce, err := cloudevents.NewClient(t, cloudevents.WithUUIDs(), cloudevents.WithTimeNow())
	if err != nil {
		log.Fatalf("failed to create cloudevents client, %s", err.Error())
	}

	log.Printf("Server starting on port 8080\n")
	if err := ce.StartReceiver(context.Background(), c.CeHandler); err != nil {
		log.Fatalf("failed to start cloudevent receiver, %s", err.Error())
	}
}
