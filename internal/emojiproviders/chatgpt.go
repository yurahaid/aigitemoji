package emojiproviders

import (
	"context"
	"encoding/json"
	"fmt"

	"aigitemoji/pkg/openai"
)

type ChatGpt struct {
	openaiClient *openai.Client
}

func NewChatGpt(openapiClient *openai.Client) *ChatGpt {
	return &ChatGpt{openaiClient: openapiClient}
}

func (c *ChatGpt) Emoji(ctx context.Context, message string) (string, error) {
	messages := []openai.Message{
		openai.NewMessage(openai.SystemRole, "You are a developers assistant designed to output JSON. It should have the fields 'emoji'"),
		openai.NewMessage(
			openai.UserRole,
			fmt.Sprintf("Mandatory сhoose the right emoji from this link https://gist.github.com/rxaviers/7360908#file-gistfile1-md based on the commit message: %s", message),
		),
	}

	response, err := c.openaiClient.Completions(ctx, messages)
	if err != nil {
		return "", fmt.Errorf("openai client completions: %v", err)
	}

	emoji := struct {
		Emoji string `json:"emoji"`
	}{}

	if err := json.Unmarshal([]byte(response.Choices[0].Message.Content), &emoji); err != nil {
		return "", fmt.Errorf("json unmarshal: %w", err)
	}

	return emoji.Emoji, nil
}
