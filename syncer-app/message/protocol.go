package message

import "encoding/json"

type ServerMessage struct{}

type StdMessage struct {
	Startline int      `json:"start_line"`
	Endline   int      `json:"end_line"`
	Text      []string `json:"text"`
}

func (msg *StdMessage) Serialize() string {
	str, _ := json.Marshal(msg)
	return string(str)
}

func Deserialize(str string) *StdMessage {
	msg := StdMessage{Text: make([]string, 0)}
	_ = json.Unmarshal([]byte(str), &msg)
	return &msg
}
