package ollama

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const (
	url   = "http://localhost:11434/api/generate"
	model = "mistral"
)

type request struct {
	Model  string `json:"model"`
	Prompt string `json:"prompt"`
}

type response struct {
	Response string `json:"response"`
}

func Query(prompt string) (string, error) {
	reqBody, err := json.Marshal(request{
		Model:  model,
		Prompt: prompt,
	})
	if err != nil {
		return "", fmt.Errorf("error marshaling request: %w", err)
	}

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(reqBody))
	if err != nil {
		return "", fmt.Errorf("error sending request to Ollama: %w", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("error reading response body: %w", err)
	}

	var ollamaResp response
	if err := json.Unmarshal(body, &ollamaResp); err != nil {
		return "", fmt.Errorf("error unmarshaling response: %w", err)
	}

	return ollamaResp.Response, nil
}
