package message

const (
	UserMessage = iota + 1
)

type Message struct {
	MessageType int    `json:"message_type,omitempty"`
	Data        []byte `json:"data,omitempty"`
}
