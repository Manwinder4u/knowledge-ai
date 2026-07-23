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

type generateOptions struct {
	Temperature float32 `json:"temperature"`
	NumPredict  int     `json:"num_predict"`
}

type generateRequest struct {
	Model  string `json:"model"`
	Prompt string `json:"prompt"`
	Stream bool   `json:"stream"`
	//Options generateOptions `json:"options"`
}

type generateResponse struct {
	Response string `json:"response"`
}

func (c *OllamaClient) Generate(ctx context.Context, prompt string) (string, error) {

	request := generateRequest{
		Model:  c.model,
		Prompt: prompt,
		Stream: false,
		// Options: generateOptions{
		// 	Temperature: 0.1,
		// 	NumPredict:  128,
		// },
	}

	body, err := json.Marshal(request)
	if err != nil {
		return "", err
	}

	req, err := http.NewRequestWithContext(
		ctx,
		http.MethodPost,
		c.baseURL+"/api/generate",
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

	if resp.StatusCode != http.StatusOK {

		data, _ := io.ReadAll(resp.Body)

		return "", fmt.Errorf(
			"ollama returned %d: %s",
			resp.StatusCode,
			string(data),
		)
	}

	var response generateResponse

	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return "", err
	}

	return response.Response, nil
}
