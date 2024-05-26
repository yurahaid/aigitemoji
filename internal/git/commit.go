package git

import (
	"fmt"

	"github.com/go-git/go-git/v5"
)

func (g *Git) Commit(message string, amend bool) (string, error) {
	hash, err := g.wt.Commit(message, &git.CommitOptions{
		Amend: amend,
	})
	if err != nil {
		return "", fmt.Errorf("fail to commit message: %w", err)
	}

	return hash.String(), nil
}
