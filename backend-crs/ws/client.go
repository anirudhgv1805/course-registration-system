package ws

import (
    "log"

    "github.com/gorilla/websocket"
)

type Client struct {
    ID           string
    DepartmentID string
    Conn         *websocket.Conn
    Send         chan []byte
    Hub          *Hub
}

func (c *Client) ReadPump() {
    defer func() {
        c.Hub.unregister <- c
        c.Conn.Close()
    }()
    for {
        _, _, err := c.Conn.ReadMessage()
        if err != nil {
            log.Println("read error:", err)
            break
        }
    }
}

func (c *Client) WritePump() {
    defer c.Conn.Close()
    for {
        msg, ok := <-c.Send
        if !ok {
            c.Conn.WriteMessage(websocket.CloseMessage, []byte{})
            return
        }
        c.Conn.WriteMessage(websocket.TextMessage, msg)
    }
}
