package ai_api

import (
	"bytes"
	"encoding/json"
	"net/http"
)

type AIAPI struct {
	APIURL string
}

func NewAIAPI() *AIAPI {
	return &AIAPI{
		APIURL: "https://wgpt-production.up.railway.app/v1/chat/completions",
	}
}

func (a *AIAPI) GetChatCompletion(messages []map[string]string) (string, error) {
	// Membuat payload dari data pesan
	payload, err := json.Marshal(map[string]interface{}{
		"model":    "gpt-3.5-turbo",
		"messages": messages,
	})
	if err != nil {
		return "", err
	}

	// Mengirim permintaan HTTP
	resp, err := http.Post(a.APIURL, "application/json", bytes.NewBuffer(payload))
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	// Membaca respons
	var response ChatCompletionResponse
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return "", err
	}

	// Memeriksa apakah ada pilihan yang tersedia
	if len(response.Choices) == 0 {
		return "", nil
	}

	// Mengambil content dari pilihan pertama
	content := response.Choices[0].Message.Content

	return content, nil
}
