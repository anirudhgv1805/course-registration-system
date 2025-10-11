package ws

import (
    "net/http"

    "github.com/gin-gonic/gin"
    "github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
    CheckOrigin: func(r *http.Request) bool {
        return true
    },
}

func ServeWS(hub *Hub) gin.HandlerFunc {
    return func(c *gin.Context) {
        // Mocked authentication — replace with real
        userID := c.Query("userID")
        departmentID := c.Query("departmentID")

        conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
        if err != nil {
            return
        }

        client := &Client{
            ID:           userID,
            DepartmentID: departmentID,
            Conn:         conn,
            Send:         make(chan []byte, 256),
            Hub:          hub,
        }

        hub.register <- client

        go client.WritePump()
        go client.ReadPump()
    }
}
