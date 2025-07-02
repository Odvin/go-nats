package main

import (
	"fmt"
	"github.com/nats-io/nats.go"
	"log"
	"math/rand"
	"time"
)

func main() {
	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		log.Fatalf("can't connect to NATS: %v", err)
	}
	defer nc.Close()

	gameId := rand.Intn(10000)
	count := 0

	go receiveClientMessage(nc, gameId)

	for {
		go sendServerMessage(nc, gameId, count)
		go sendClientMessage(nc, gameId, count)
		time.Sleep(2 * time.Second)
		count++
	}

}

func sendServerMessage(nc *nats.Conn, gameId, count int) {
	message := fmt.Sprintf("Message Alpha(%d) game: %d", gameId, count)

	err := nc.Publish("game.alpha.message", []byte(message))
	if err != nil {
		log.Fatalf("can't send message to NATS subject 'game.alpha.message': %v", err)
	}

	log.Printf("Alpha(%d) game sent message %d", gameId, count)
}

func sendClientMessage(nc *nats.Conn, gameId, count int) {
	message := fmt.Sprintf("Event Alpha(%d) game: %d", gameId, count)

	err := nc.Publish("game.alpha.client-event", []byte(message))
	if err != nil {
		log.Fatalf("can't send event to NATS subject 'game.alpha.client-event': %v", err)
	}

	log.Printf("Alpha(%d) game sent event %d", gameId, count)
}

func receiveClientMessage(nc *nats.Conn, gameId int) {
	_, err := nc.Subscribe("game.alpha.client-event", func(m *nats.Msg) {
		fmt.Printf("Alpha(%d) game got an event: %s\n", gameId, string(m.Data))
	})
	if err != nil {
		log.Fatalf("can't subscribe to NATS subject 'game.alpha.client-event': %v", err)
	}
}
