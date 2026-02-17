package network

import "encoding/json"

type MessageType string

const (
	MsgJoin   MessageType = "join"
	MsgStart  MessageType = "start"
	MsgUpdate MessageType = "update"
)

type Message struct {
	Type    MessageType     `json:"type"`
	Payload json.RawMessage `json:"payload"`
}

type JoinPayload struct {
	PlayerName string `json:"player_name"`
}

type StartPayload struct {
	Text      string `json:"text"`
	Countdown int    `json:"countdown"`
}

type UpdatePayload struct {
	PlayerName string  `json:"player_name"`
	Progress   float64 `json:"progress"`
	WPM        int     `json:"wpm"`
	Completed  bool    `json:"completed"`
}
