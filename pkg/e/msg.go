package e

var codeMsg = map[Code]string{
	WebsocketSuccessMessage: "解析content内容信息",
	WebsocketSuccecss:       "发送信息，请求历史记录操作成功",
	WebsocketEnd:            "请求历史记录，但没有更多记录了",
	WebsocketOnlineReply:    "针对回复信息在线应答成功",
	WebsocketOffineReply:    "针对回复信息离线应答成功",
	WebsocketLimit:          "请求受到限制",
}
