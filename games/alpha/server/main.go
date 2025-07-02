package main

import (
	"fmt"
	"github.com/nats-io/nats.go"
	"log"
	"time"
)

func main() {
	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		log.Fatalf("can't connect to NATS: %v", err)
	}
	defer nc.Close()

	_, err = nc.Subscribe("game.alpha.*", func(m *nats.Msg) {
		fmt.Printf("Alpha game server got a message: %s\n", string(m.Data))
	})
	if err != nil {
		log.Fatalf("can't subscribe to NATS subject 'game.alpha.*': %v", err)
	}

	time.Sleep(1 * time.Hour)
}
