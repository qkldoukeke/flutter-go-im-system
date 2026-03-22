package main

import (
	"fmt"
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

type Message struct {
	From    string `json:"from"`
	To      string `json:"to"`
	Content string `json:"content"`
	Time    int64  `json:"time"`
}

var wsUpgrader = websocket.Upgrader{CheckOrigin: func(r *http.Request) bool { return true }}

var (
	connections = make(map[string]*websocket.Conn) // 用户id -> 连接
	lock        sync.Mutex
)

// WebSocket 聊天核心逻辑，包含用户上线/下线及点对点消息
func main() {
	r := gin.Default()
	r.GET("/ws/:uid", func(c *gin.Context) {
		uid := c.Param("uid")
		conn, _ := wsUpgrader.Upgrade(c.Writer, c.Request, nil)
		lock.Lock()
		connections[uid] = conn
		lock.Unlock()
		fmt.Println("用户上线", uid)
		defer func() {
			conn.Close()
			lock.Lock()
			delete(connections, uid)
			lock.Unlock()
		}()
		for {
			var msg Message
			err := conn.ReadJSON(&msg)
			if err != nil {
				break
			}
			// 发送给目标用户
			lock.Lock()
			if toConn, ok := connections[msg.To]; ok {
				toConn.WriteJSON(msg)
			}
			lock.Unlock()
		}
	})
	r.Run(":8082")
}