package models

type Msg struct {
	ChatID int64  `json:"chat_id"`
	Text   string `json:"text"`
}
