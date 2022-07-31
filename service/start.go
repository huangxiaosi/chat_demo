package service

import (
	"chat-demo/conf"
	"chat-demo/pkg/e"
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
)

func (manager *ClientManager) Start() {
	for {
		fmt.Println("--------监听管道通讯-----------")
		select {
		case conn := <-Manager.Register:
			fmt.Printf("有新链接：%s", conn.ID)
			Manager.Clients[conn.ID] = conn //把连接放到管理上
			replyMsg := ReplyMsg{
				Code:    e.WebsocketSuccecss,
				Content: "已经连接到服务器",
			}
			msg, _ := json.Marshal(replyMsg)
			_ = conn.Socket.WriteMessage(websocket.TextMessage, msg)
		case conn := <-manager.Unregister:
			fmt.Println("连接失败%s", conn.ID)
			if _, ok := manager.Clients[conn.ID]; ok {
				replyMsg := &ReplyMsg{
					Code:    e.WebsocketEnd,
					Content: "连接中断",
				}
				msg, _ := json.Marshal(replyMsg)
				_ = conn.Socket.WriteMessage(websocket.TextMessage, msg)
				close(conn.Send)
				delete(manager.Clients, conn.ID)
			}
		case broadcase := <-Manager.Broadcase: //1->2
			message := broadcase.Message
			sendId := broadcase.Client.SendID //2->1
			flag := false                     //默认对方不在线
			for id, conn := range Manager.Clients {
				if id != sendId {
					continue
				}
				select {
				case conn.Send <- message:
					flag = true
				default:
					close(conn.Send)
					delete(Manager.Clients, conn.ID)
				}
				id := broadcase.Client.ID //1->2
				if flag {
					replyMsg := &ReplyMsg{
						Code:    e.WebsocketOnlineReply,
						Content: "对方在线应答",
					}
					msg, _ := json.Marshal(replyMsg)
					_ = broadcase.Client.Socket.WriteMessage(websocket.TextMessage, msg)

					err := InsertMsg(conf.MongoDBName, id, string(message), 1, int64(3*month)) //对方在线，状态1，默认消息已读
					if err != nil {
						fmt.Println("InsertOne Err:", err)
					} else {
						fmt.Println("mongo插入成功")
					}
				}
			}

		}
	}
}
