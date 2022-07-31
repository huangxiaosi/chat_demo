package e

type Code int

const (
	WebsocketSuccessMessage = 50001
	WebsocketSuccecss       = 50002
	WebsocketEnd            = 50003
	WebsocketOnlineReply    = 50004
	WebsocketOffineReply    = 50005
	WebsocketLimit          = 50006
)

func (c Code) Msg() string {
	return codeMsg[c]
}
