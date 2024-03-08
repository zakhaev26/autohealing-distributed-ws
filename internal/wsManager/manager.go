package wsManager

import (
	"log"
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
)

var (
	websocketUpgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}
)

type Manager struct {
	clients ClientList

	locker sync.RWMutex // for editing the clientList safely
}

func NewManager() *Manager {
	return &Manager{
		clients: make(ClientList),
	}
}

func (m *Manager) addClient(client *Client) {
	m.locker.Lock() //lock so we cant manipulate
	defer m.locker.Unlock()
	m.clients[client] = true
}

// func (m *Manager) removeClient(client *Client) {
// 	m.locker.Lock()
// 	defer m.locker.Unlock()

// 	if _, ok := m.clients[client]; ok {
// 		client.conn.Close()
// 		delete(m.clients, client)
// 	}
// }

func (m *Manager) ServeWS(w http.ResponseWriter, r *http.Request) {
	log.Println("New Connection")
	conn, err := websocketUpgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	client := NewClient(conn, m)
	m.addClient(client)

	go client.conn.ReadMessage()
}
