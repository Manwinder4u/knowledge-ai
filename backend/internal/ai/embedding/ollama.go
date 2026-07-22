package embedding

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type OllamaEmbedder struct {
	baseURL string
	model   string
}

// Constructor
func NewOllamaEmbedder(baseURL string, model string) *OllamaEmbedder {
	return &OllamaEmbedder{
		baseURL: baseURL,
		model:   model,
	}
}

type embeddingRequest struct {
	Model  string `json:"model"`
	Prompt string `json:"prompt"`
}

type embeddingResponse struct {
	Embedding []float32 `json:"embedding"`
}

func (e *OllamaEmbedder) Embed(ctx context.Context, text string) ([]float32, error) {

	request := embeddingRequest{
		Model:  e.model,
		Prompt: text,
	}

	fmt.Println("Request: ", request)
	fmt.Println("Embedder:", e.baseURL)
	body, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(
		ctx,
		http.MethodPost,
		e.baseURL+"/api/embeddings",
		bytes.NewReader(body),
	)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {

		data, _ := io.ReadAll(resp.Body)

		return nil, fmt.Errorf(
			"ollama returned %d: %s",
			resp.StatusCode,
			string(data),
		)
	}

	var response embeddingResponse

	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return nil, err
	}

	return response.Embedding, nil
}
