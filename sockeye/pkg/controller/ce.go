package controller

import (
	"encoding/json"
	"fmt"
	cloudevents "github.com/cloudevents/sdk-go"
)

func (c *Controller) CeHandler(event cloudevents.Event) {
	fmt.Println("got", event.String())

	b, err := json.Marshal(event)
	if err != nil {
		fmt.Println("err", err)
		return
	}

	// TODO: cloudevents needs a websocket transport.

	manager.broadcast <- string(b)

	return
}
