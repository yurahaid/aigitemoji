package openai

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Model string

const (
	Model35turbo Model = "gpt-3.5-turbo"
)

type Role string

const (
	SystemRole Role = "system"
	UserRole   Role = "user"
)

type Client struct {
	httpClient *http.Client
	url        string
	apiToken   string
	model      Model
}

type Message struct {
	Role    Role   `json:"role"`
	Content string `json:"content"`
}

func NewMessage(role Role, content string) Message {
	return Message{Role: role, Content: content}
}

type Response struct {
	ID      string `json:"id"`
	Object  string `json:"object"`
	Created int64  `json:"created"`
	Model   string `json:"model"`
	Choices []struct {
		Message      Message `json:"message"`
		FinishReason string  `json:"finish_reason"`
		Index        int     `json:"index"`
	} `json:"choices"`
}

type request struct {
	Model          Model          `json:"model"`
	Messages       []Message      `json:"messages"`
	ResponseFormat responseFormat `json:"response_format"`
}

type responseFormat struct {
	Type string `json:"type"`
}

func NewClient(httpClient *http.Client, url string, apiToken string, model Model) *Client {
	return &Client{
		httpClient: httpClient,
		url:        url,
		apiToken:   apiToken,
		model:      model,
	}
}

func (c *Client) Completions(ctx context.Context, messages []Message) (Response, error) {
	body, err := json.Marshal(request{
		Model:          c.model,
		Messages:       messages,
		ResponseFormat: responseFormat{Type: "json_object"},
	})

	if err != nil {
		return Response{}, fmt.Errorf("marshalling messages: %w", err)
	}
	req, err := http.NewRequest(http.MethodPost, fmt.Sprintf("%s/v1/chat/completions", c.url), bytes.NewBuffer(body))
	if err != nil {
		return Response{}, fmt.Errorf("creating request: %w", err)
	}

	httpResp, err := c.sendRequest(ctx, req)
	if err != nil {
		return Response{}, fmt.Errorf("doing request: %w", err)
	}
	defer httpResp.Body.Close()

	if httpResp.StatusCode != http.StatusOK {
		respBody, _ := io.ReadAll(httpResp.Body)
		return Response{}, fmt.Errorf("unexpected http status code: %d, response %s", httpResp.StatusCode, respBody)
	}

	var response Response
	if err := json.NewDecoder(httpResp.Body).Decode(&response); err != nil {
		return Response{}, fmt.Errorf("decoding response: %w", err)
	}

	return response, nil
}

func (c *Client) sendRequest(ctx context.Context, req *http.Request) (*http.Response, error) {
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.apiToken))
	req.Header.Set("Content-Type", "application/json")

	return c.httpClient.Do(req.WithContext(ctx))
}
