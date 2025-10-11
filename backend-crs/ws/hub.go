package ws

import (
	"encoding/json"
	"sync"
)

type Hub struct {
	clients    map[*Client]bool
	register   chan *Client
	unregister chan *Client
	mu         sync.RWMutex
}

func NewHub() *Hub {
	return &Hub{
		clients:    make(map[*Client]bool),
		register:   make(chan *Client),
		unregister: make(chan *Client),
	}
}

func (h *Hub) Run() {
	for {
		select {
		case client := <-h.register:
			h.mu.Lock()
			h.clients[client] = true
			h.mu.Unlock()

		case client := <-h.unregister:
			h.mu.Lock()
			if _, ok := h.clients[client]; ok {
				delete(h.clients, client)
				close(client.Send)
			}
			h.mu.Unlock()
		}
	}
}

func (h *Hub) BroadcastToDepartment(departmentID string, message any) {
	data, _ := json.Marshal(message)
	h.mu.RLock()
	defer h.mu.RUnlock()
	for c := range h.clients {
		if c.DepartmentID == departmentID {
			select {
			case c.Send <- data:
			default:
				close(c.Send)
				delete(h.clients, c)
			}
		}
	}
}
