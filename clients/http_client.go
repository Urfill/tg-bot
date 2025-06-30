package clients

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"tg-bot/contants"
	"tg-bot/models"
)

type HttpClient struct {
	client *http.Client
}

func NewHttpClient() *HttpClient {
	return &HttpClient{
		client: &http.Client{},
	}
}

func (h *HttpClient) SendMsg(msg models.Msg) (*http.Response, error) {
	sendMessageURL := fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage", contants.BotToken) // не нужно отправлять всё в открытом виде. используй tg client пакет
	msgJson, err := json.Marshal(msg)
	if err != nil {
		return nil, err
	}
	resp, err := h.client.Post(sendMessageURL, "application/json", bytes.NewBuffer(msgJson))
	if err != nil {
		return nil, err
	}

	switch resp.StatusCode {
	case http.StatusOK:
		return resp, nil
	default:
		return nil, fmt.Errorf("failed to send msg, returned code: %d", resp.StatusCode)
	}
}
