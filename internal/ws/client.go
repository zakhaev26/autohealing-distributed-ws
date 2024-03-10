package ws

import (
	"fmt"
	"log"

	"github.com/gorilla/websocket"
	"github.com/zakhaev26/pteromys/internal/kafka"
)

type Client struct {
	connection *websocket.Conn
	manager    *Manager

	egress chan []byte
}

type ClientList map[*Client]bool

func NewClient(c *websocket.Conn, m *Manager) *Client {
	return &Client{
		connection: c,
		manager:    m,
		egress:     make(chan []byte),
	}
}

func (c *Client) readMessages() {

	defer func() {
		c.manager.removeClient(c)
		log.Println("client got disconnected.")
	}()

	for {
		_,_, err := c.connection.ReadMessage()

		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("error reading message: %v", err)
			}
			return
		}
	}
}

func (c *Client) writeMessages() {
	defer func() {
		c.manager.removeClient(c)
	}()

	consumer, _ := kafka.Consumer("deez")

	go func() {
		log.Println("uhm in go thread running kafka dawg")
		for {
			select {
			case err := <-consumer.Errors():
				fmt.Println(err)
			case msg := <-consumer.Messages():
				var message []byte = msg.Value
				c.egress <- message
			}
		}

	}()

	for {
		select {
		case msg, ok := <-c.egress:
			if !ok {

				if err := c.connection.WriteMessage(websocket.CloseMessage, nil); err != nil {
					log.Println("connection closed :", err)
				}
				return
			}

			if err := c.connection.WriteMessage(websocket.TextMessage, msg); err != nil {
				log.Println("error in writing mesage to client : ", err)
				return
			}
			log.Println("sent message ; clientList length =", len(c.manager.clientList))
		}
	}
}
