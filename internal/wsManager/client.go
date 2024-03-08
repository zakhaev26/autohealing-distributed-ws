package wsManager

import "github.com/gorilla/websocket"

type Client struct {
	conn    *websocket.Conn
	manager *Manager
}

func NewClient(conn *websocket.Conn, manager *Manager) *Client {
	return &Client{
		conn:    conn,
		manager: manager,
	}
}

type ClientList map[*Client]bool
