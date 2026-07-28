package llm

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type OllamaClient struct {
	baseURL string
	model   string
}

func NewOllamaClient(baseURL string, model string) *OllamaClient {
	return &OllamaClient{
		baseURL: baseURL,
		model:   model,
	}
}

type chatOptions struct {
	Temperature float32 `json:"temperature"`
	NumPredict  int     `json:"num_predict"`
}
type chatRequest struct {
	Model    string        `json:"model"`
	Messages []chatMessage `json:"messages"`
	Stream   bool          `json:"stream"`
	Options  chatOptions   `json:"options"`
}
type chatMessage struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}
type chatResponse struct {
	Message struct {
		Role    string `json:"role"`
		Content string `json:"content"`
	} `json:"message"`

	Done bool `json:"done"`
}

func (c *OllamaClient) Generate(ctx context.Context, messages []Message) (string, error) {
	ollamaMessages := make([]chatMessage, 0, len(messages))

	for _, message := range messages {
		ollamaMessages = append(
			ollamaMessages,
			chatMessage{
				Role:    string(message.Role),
				Content: message.Content,
			},
		)
	}

	request := chatRequest{
		Model:    c.model,
		Messages: ollamaMessages,
		Stream:   false,
		Options: chatOptions{
			Temperature: 0.1,
			NumPredict:  512,
		},
	}

	body, err := json.Marshal(request)
	if err != nil {
		return "", err
	}

	fmt.Printf("%+v\n", ollamaMessages)
	req, err := http.NewRequestWithContext(
		ctx,
		http.MethodPost,
		c.baseURL+"/api/chat",
		bytes.NewReader(body),
	)
	if err != nil {
		return "", err
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	fmt.Println("========== OLLAMA RESPONSE ==========")
	fmt.Println(string(data))
	fmt.Println("=====================================")
	// if resp.StatusCode != http.StatusOK {

	// 	// return "", fmt.Errorf(
	// 	// 	"ollama returned %d: %s",
	// 	// 	resp.StatusCode,
	// 	// 	string(data),
	// 	// )
	// }

	var response chatResponse

	if err := json.Unmarshal(data, &response); err != nil {
		return "", err
	}

	fmt.Printf("%+v\n", response)

	return response.Message.Content, nil
}
