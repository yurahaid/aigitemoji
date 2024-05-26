package git

import (
	"fmt"
	"os"

	"github.com/go-git/go-git/v5"
)

type Git struct {
	repo *git.Repository
	wt   *git.Worktree
}

func SetupGit() (*Git, error) {
	wd, err := os.Getwd()
	if err != nil {
		return nil, fmt.Errorf("fail to get working directory: %w", err)
	}

	r, err := git.PlainOpen(wd)
	if err != nil {
		return nil, fmt.Errorf("fail to open git repository: %w", err)
	}

	wt, err := r.Worktree()
	if err != nil {
		return nil, fmt.Errorf("fail to get working tree: %w", err)
	}

	return &Git{
		repo: r,
		wt:   wt,
	}, nil
}
