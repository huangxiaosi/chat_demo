package service

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"net/http"
)

const (
	month = 60 * 60 * 24 * 30
)

type SendMsg struct {
	Type    int    `json:"type"`
	Content string `json:"content"`
}

type ReplyMsg struct {
	From    string `json:"from"`
	Code    int    `json:"code"`
	Content string `json:"content"`
}

type Client struct {
	ID     string
	SendID string
	Socket *websocket.Conn
	Send   chan []byte
}

//广播类
type Broadcase struct {
	Client  *Client
	Message []byte
	Type    int
}

//用户管理
type ClientManager struct {
	Clients    map[string]*Client
	Broadcase  chan *Broadcase
	Reply      chan *Client
	Register   chan *Client
	Unregister chan *Client
}

type Message struct {
	Sender    string `json:"sender",omitempty`
	Recipient string `json:"recipient",omitempty`
	Content   string `json:"content",omitempty`
}

var Manager = ClientManager{
	Clients:    make(map[string]*Client),
	Broadcase:  make(chan *Broadcase),
	Register:   make(chan *Client),
	Reply:      make(chan *Client),
	Unregister: make(chan *Client),
}

func CreateID(uid, toUid string) string {
	return uid + "->" + toUid
}

func Handler(c *gin.Context) {
	uid := c.Query("id")
	toUid := c.Query("toUid")
	conn, err := (&websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		}}).Upgrade(c.Writer, c.Request, nil) //升级ws协议
	if err != nil {
		http.NotFound(c.Writer, c.Request)
		return
	}

	//创建一个用户实例
	client := &Client{
		ID:     CreateID(uid, toUid), // 1->2
		SendID: CreateID(toUid, uid), //2->1
		Socket: conn,
		Send:   make(chan []byte),
	}

	//注册用户
	Manager.Register <- client
	go client.Read()
	go client.Write()
}

func (c *Client) Read {
	
}

func (c *Client) Write {

}
