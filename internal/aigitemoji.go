package internal

import (
	"context"
	"fmt"

	"github.com/yurahaid/aigitemoji/internal/git"
)

type EmojiProvider interface {
	Emoji(ctx context.Context, message string) (string, error)
}

type AIGitEmoji struct {
	emojiProvider EmojiProvider
	git           *git.Git
}

func NewAIGitEmoji(emojiProvider EmojiProvider, git *git.Git) *AIGitEmoji {
	return &AIGitEmoji{emojiProvider: emojiProvider, git: git}
}

func (a *AIGitEmoji) Commit(ctx context.Context, message string, amend bool) (commit string, hash string, err error) {
	emoji, err := a.emojiProvider.Emoji(ctx, message)
	if err != nil {
		return "", "", fmt.Errorf("error getting emoji: %w", err)
	}

	emojiCommit := fmt.Sprintf("%s %s", emoji, message)

	if hash, err = a.git.Commit(emojiCommit, amend); err != nil {
		return "", "", fmt.Errorf("error getting emoji: %w", err)
	}

	return emojiCommit, hash, nil
}
